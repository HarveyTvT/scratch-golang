package main

import "net"

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range interfaces {
		println("Interface Name:", iface.Name)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			println("Address:", addr.String())
		}
	}
}
