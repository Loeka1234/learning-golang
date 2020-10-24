package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	port := flag.Int("p", 1234, "The port the server should be running on." )
	flag.Parse()

	fmt.Printf("Starting TCP server on PORT: %v\n", *port)

	PORT := ":" + strconv.Itoa(*port)
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("Error while starting server: ", err)
	}

	defer listener.Close()

	c, err := listener.Accept()
	if err != nil {
		log.Fatal("Error while accepting connection: ", err)
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(netData) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		log.Print("Data from request: ", netData)
		t := time.Now()
		myTime := "Today's time: " +  t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}