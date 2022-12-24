package app

import (
	"identity-service/internal/transport"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/health", transport.HandleHealth)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}
