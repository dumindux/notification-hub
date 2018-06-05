package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"notification-hub/utils"
	"notification-hub/monitor"
	"strconv"
)

var upgrader = websocket.Upgrader{
}

func main() {
	log.Println("starting the notification hub")
	config := utils.LoadConfig()
	dataChannel := make(chan []byte)
	go utils.DataExporter(&config.InfluxDB, config.Application.Title, dataChannel)
	startServer(config.Application.Port, config.Application.Key, config.Application.Cert, config.Application.Interval, dataChannel)
}

func startServer(port int, key string, cert string, interval int,  dataChannel chan []byte) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("client connected from " + r.Host);
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go monitor.ReadFromDeamon(conn, interval, dataChannel)
	})
	log.Println("server listening...")
	err := http.ListenAndServeTLS(":" + strconv.Itoa(port), cert, key,nil)
	if err != nil {
		log.Fatal("listen and serve:", err)
	}
}