package handlers

import (
	"net/http"
	"path/filepath"

)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("frontend", "index.html")
	http.ServeFile(w, r, filePath)
}