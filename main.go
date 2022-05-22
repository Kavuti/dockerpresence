package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
)

var port int

func init() {

	flag.IntVar(&port, "port", 9010, "The port Dockerpresence will use")
}

func main() {
	fmt.Printf("Starting Dockerpresence server on port %d", port)
	stringedPort := ":" + strconv.Itoa(port)
	server, err := net.Listen("tcp4", stringedPort)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println("Error acceppting a new connection")
			panic(err)
		}
		go handleConnection(client)
	}
}
s
func handleConnection(client net.Conn) {
	reader := bufio.NewReader(client)
	for {
		singleByte, err := reader.ReadByte()
		if err != nil {
			fmt.Println("Finished reading bytes")
			break
		}
		fmt.Printf("%s", string(singleByte))
	}
	client.Close()
}
