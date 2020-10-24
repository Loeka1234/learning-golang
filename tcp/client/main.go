package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	host := flag.String("host", "127.0.0.1", "The destination hostname.")
	port := flag.Int("port", 1234, "The port to use.")
	flag.Parse()

	CONNECT := *host + ":" + strconv.Itoa(*port)
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Request: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprintf(c, text + "\n")
		if err != nil {
			panic(err)
		}

		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("Response: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}