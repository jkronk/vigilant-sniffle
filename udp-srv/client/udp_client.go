package main

import (
	"net"
	"fmt"
	"strconv"
)

const serverAddress = "0.0.0.0:3000"

func main() {
	fmt.Println("UDP Client Shell")
	fmt.Println("Sending packets to the server address")

	serverAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		panic(err)
	}

	localAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3001")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil { 
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < 100; i++ {
		msg := "test message"
		buf := []byte(msg)
		
		fmt.Printf("sending message #" + strconv.Itoa(i) + "\r\n")
		_, err := conn.Write(buf)
		if err != nil { 
			fmt.Println(msg, err)
		}
	}

	select {}
}