package main

import (
	"fmt"
	"go-chatroom/handlers"
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.HomeHandler)
	fmt.Print("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)

}