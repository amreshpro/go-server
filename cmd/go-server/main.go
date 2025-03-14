package main

import (
	"log"
	"net/http"
	"github.com/amreshpro/go-server/internal/router"
)

func main() {
	r := router.NewRouter()

	log.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
