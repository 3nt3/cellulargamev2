package main

import (
	"cellulargamev2/funcs"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// idk ... structs i think
type ClientRequest struct {
	Type string `json:"type"`// [spawnFood, getFood, initCell, updateSize, (delall), eat, getCells]
	Data string `json:"data"`
}

type ClientResponse struct {
	Data []byte
}

// vars
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}



func main() {
	http.HandleFunc("/", handler)
	go log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		creq := &ClientRequest{}
		err := conn.ReadJSON(creq)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("Message from client: %v", creq)

		cresp := &ClientResponse{}
		switch creq.Type {
		case "spawnFood":
			cresp.Data, _ = json.Marshal(funcs.SpawnFood())

		case "getFood":
			cresp.Data, _ = json.Marshal(funcs.GetFood())

		case "initCell":
			var data map[string]string
			_ = json.Unmarshal([]byte(creq.Data), &data)
			log.Println(data)
			name := data["name"]
			//log.Println(reflect.TypeOf(data))
			cresp.Data, _ = json.Marshal(funcs.InitCell(name))

		case "updateSize":
			var data map[string]int
			_ = json.Unmarshal([]byte(creq.Data), &data)
			id := data["id"]
			size := data["size"]
			cresp.Data, _ = json.Marshal(funcs.ChangeSize(id, size))

		case "delall":
			funcs.Delall()

		case "eat":
			var data map[string]int
			_ = json.Unmarshal([]byte(creq.Data), &data)
			id := data["id"]
			mealId := data["mealId"]

			cresp.Data, _ = json.Marshal(funcs.Eat(id, mealId))

		case "getCells":
			cresp.Data, _ = json.Marshal(funcs.GetCells())
		}
		_ = conn.WriteJSON(cresp)
	}
}
