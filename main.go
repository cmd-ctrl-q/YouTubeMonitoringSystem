package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cmd-ctrl-q/YouTubeMonitoringSystem/websocket"
	"github.com/cmd-ctrl-q/YouTubeMonitoringSystem/youtube"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("Youtube Subscriber Monitor")

	item, err := youtube.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", item)

	setupRoutes()
}
