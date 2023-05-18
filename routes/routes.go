package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func HandlePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		return
	}
	fmt.Println(message)
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		return
	}
	
}