package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	r := http.NewServeMux()
	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	fmt.Println("Server running on port:", port)
	fmt.Println("DB User:", dbUser)
	fmt.Println("DB Password:", dbPass)
	log.Println("Starting server at :8080")
	log.Println("AppName: "+os.Getenv("APP_NAME"))
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server failed:", err)
		}
	
	
}
