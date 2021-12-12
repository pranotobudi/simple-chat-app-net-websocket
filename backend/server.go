package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	// var err error

	// for {
	// 	var reply string

	// 	if err = websocket.Message.Receive(ws, &reply); err != nil {
	// 		fmt.Println("Can't receive")
	// 		break
	// 	}

	// 	fmt.Println("Received back from client: " + reply)

	// 	msg := "Received:  " + reply
	// 	fmt.Println("Sending to client: " + msg)

	// 	if err = websocket.Message.Send(ws, msg); err != nil {
	// 		fmt.Println("Can't send")
	// 		break
	// 	}
	// }

	// receive text frame
	var message string
	websocket.Message.Receive(ws, &message)

	// send text frame
	message = "hello"
	websocket.Message.Send(ws, message)

	// receive binary frame
	var data []byte
	websocket.Message.Receive(ws, &data)

	// send binary frame
	data = []byte{0, 1, 2}
	websocket.Message.Send(ws, data)

}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	log.Println("Server Running on port 1234...")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
