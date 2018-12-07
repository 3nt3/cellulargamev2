package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ws", func (w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			log.Printf("Error creating websocket connection: %v\n", err)
			return
		}
		ws.On("message", func(e *Event) {
			log.Printf("Message received: %s", e.Data.(string))
		})
	})

	http.ListenAndServe(":8000", nil)

}