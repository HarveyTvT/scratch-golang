package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var (
	upgrader  = websocket.Upgrader{}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
)

// server
func handleConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	clients[ws] = true
	for {
		t, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(clients, ws)
			break
		}
		if t != websocket.TextMessage {
			continue
		}

		broadcast <- string(msg)
	}
}

func handleMessage() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func start() {
	http.HandleFunc("/ws", handleConnection)
	go handleMessage()
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

// client
func connect() {
	c, _, err := websocket.DefaultDialer.Dial("ws://:8080/ws", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	done := make(chan struct{})

	// read message
	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("recv: %s\n", string(message))
		}
	}()

	// write message
	go func() {
		defer close(done)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := scanner.Text()
			if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("send: %s\n", msg)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			fmt.Println("interrupt")
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "start":
		start()
	case "connect":
		connect()
	}

}
