package main

import (
"fmt"
"net"
"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "1883"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	for {
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		fmt.Printf("Received message, size: %d %d\n", len(buf), reqLen)
		for i := 0; i <= reqLen; i++{
			fmt.Printf("%x ", buf[i])
		}
		fmt.Printf("\n")

		//Start processing MQTT packets
		if (buf[0] == 0x10) {
			fmt.Printf("MQTT Init Conex\n")
			mesg := make([]byte, 4)
			mesg[0] = 0x20
			mesg[1] = 0x02
			mesg[2] = 0x00
			mesg[3] = 0x00
			conn.Write(mesg)
		} else if (buf[0] == 0x30) {
			fmt.Printf("MQTT Publish \n")
		} else {
			fmt.Printf("Unkown Type: %x", buf[0])
		}
	}


	// Send a response back to person contacting us.
	//conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	//conn.Close()
}