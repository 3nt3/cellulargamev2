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
	Type string `json:"type"`
	Data string `json:"data"`
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
		_ = conn.ReadJSON(creq)
		log.Printf("Message from client: %v", creq)

		cresp := &ClientResponse{}
		switch creq.Type {
		case "spawnFood":
			data, _ := json.Marshal(funcs.SpawnFood())
			cresp.Data = string(data)

		case "getFood":
			data, _ := json.Marshal(funcs.GetFood())
			cresp.Data = string(data)

		case "initCell":
			var data map[string]string
			_ = json.Unmarshal([]byte(creq.Data), &data)
			log.Println(data)
			name := data["name"]
			//log.Println(reflect.TypeOf(data))
			jsonData, _ := json.Marshal(funcs.InitCell(name))
			cresp.Data = string(jsonData)

		case "updateSize":
			var data map[string]int
			_ = json.Unmarshal([]byte(creq.Data), &data)
			id := data["id"]
			size := data["size"]

			jsonData, _ := json.Marshal(funcs.ChangeSize(id, size))
			cresp.Data = string(jsonData)

		case "delall":
			funcs.Delall()

		case "eat":
			var data map[string]int
			_ = json.Unmarshal([]byte(creq.Data), &data)
			id := data["id"]
			mealId := data["mealId"]

			jsonData, _ := json.Marshal(funcs.Eat(id, mealId))
			cresp.Data = string(jsonData)

		default:
			data, _ := json.Marshal(funcs.GetCells())
			cresp.Data = string(data)
			cresp.Type = "status"
			_ = conn.WriteJSON(cresp)
			continue
		}
		conn.WriteJSON(cresp)
		cresp.Type = creq.Type
	}
}
