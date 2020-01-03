package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("UDP Server Listner Shell")
	fmt.Println("---------------------")
	
	conn, err := net.ListenUDP("udp", &net.UDPAddr {
		Port: 3000,
		IP: net.ParseIP("0.0.0.0"),
	})

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 1024)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:rlen]))
		fmt.Printf("rx: %s from %s\r\n", data, remote)
	}
}