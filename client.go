package main

import (
	"bufio"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"os"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:12345/echo"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for {
			var x string
			err = websocket.Message.Receive(ws, &x)
			if err != nil {
				fmt.Printf("error in receive %s\n", err.Error())
			}

			fmt.Printf("Recieved: %s\n", x)
		}
	}()

	for {

		bio := bufio.NewReader(os.Stdin)
		line, _, _ := bio.ReadLine()

		// fmt.Printf("Sending: %s\n", line)
		err = websocket.Message.Send(ws, line)
		if err != nil {
			fmt.Printf("error in send %s\n", err.Error())
		}
	}
}
