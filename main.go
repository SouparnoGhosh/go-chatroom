package main

import (
	"fmt"
	"go-chatroom/handlers"
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	router := chi.NewRouter()

	router.Get("/", handlers.HomeHandler)
	router.Get("/ws", handlers.WSHandler)
	fmt.Print("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)

}