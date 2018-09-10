package main

import (
	"log"
	"net/http"

	"blog/router"
)

func main() {
	router.Routes()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println("Server Start Failed")
	} else {
		log.Println("Listening on 0.0.0.0:8090")
	}
}
