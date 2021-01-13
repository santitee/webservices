package product

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func productSocket( ws *websocket.Conn ) {
	go func(c *websocket.Conn) {
		for {
			var msg message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("recieve message %s\n", msg.Data)
		}
	}(ws)

	for {
		products, err := GetTopTenProducts()
		if err != nil {
			log.Println(err)
			break
		}
		if err := websocket.JSON.Send(ws, products); err != nil {
			log.Println(err)
			break
		}
		time.Sleep(10 * time.Second)
	}
}
