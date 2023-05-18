package main

import (
	"go_modules/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", routes.HandlePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}

}