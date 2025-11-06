package main

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"strings"
)

//go:embed udhr
var udhr embed.FS

func main() {
	if len(os.Args) == 1 {
		items, err := udhr.ReadDir("udhr")
		if err != nil {
			println(err.Error())
			return
		}

		fmt.Println("Available Languages: ")
		for _, item := range items {
			if v := strings.Split(item.Name(), "_"); len(v) > 0 {
				fmt.Println(v[0])
			}
		}

		return
	}

	data, err := udhr.ReadFile("udhr/" + fmt.Sprintf("%s_rights.txt", os.Args[1]))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file not exist")
			return
		}
		fmt.Println("error: " + err.Error())
		return
	}

	fmt.Println(string(data))
}
