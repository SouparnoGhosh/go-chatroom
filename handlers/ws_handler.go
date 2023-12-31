package handlers

import (
	"net/http"
	"github.com/gorilla/websocket"
)

// An upgrader is needed to upgrade the HTTP connection to a TCP connection
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool) // connected clients

func WSHandler(w http.ResponseWriter, r *http.Request) {
	// Create a WS Connection from HTTP Connection
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
		// Sends the response to the client screen thru w (response writer), and the Status of Internal Server Error (500)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer conn.Close()

	// Infinite for loop. Read the message and then print it immediately.
    for {
        messageType, messageBody, err := conn.ReadMessage()
        if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

		// timestamp := time.Now().Format("06-01-02 15:04")
		// message := []byte(timestamp + " " + string(messageBody))

        // // Echo the message back to the client
        // if err := conn.WriteMessage(messageType, message); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
        //     return
        // }

        for client := range clients {
            err := client.WriteMessage(messageType, messageBody)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                client.Close()
                delete(clients, client)
            }
        }
    }
}
