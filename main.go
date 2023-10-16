package main

import (
	"fmt"
	"go-chatroom/handlers"
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/gin-gonic/gin"
)


func main() {
	router := chi.NewRouter()
	// router := gin.Default()


	router.Get("/", handlers.HomeHandler)
	router.Get("/ws", handlers.WSHandler)
	fmt.Print("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)

}