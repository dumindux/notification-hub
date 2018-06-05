package monitor

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Status struct {
	Time int64
	Tags map[string]string
	Fields map[string]interface{}
}

func ReadFromDeamon(connection *websocket.Conn, interval int, dataChannel chan []byte) {
	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Println("status message received from " + connection.RemoteAddr().String())
		dataChannel <- message
		time.Sleep(time.Millisecond * time.Duration(interval))
	}
}
