package main

import (
	"go_modules/routes"
	"log"
	"net/http"
)

func main() {
   http.HandleFunc("/login",routes.Login)
   http.HandleFunc("/home",routes.Home)
    log.Fatal(http.ListenAndServe(":8080",nil))
}