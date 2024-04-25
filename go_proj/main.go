package main

import (
	"go_proj/rest"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/messages/", handleRequests)
	http.HandleFunc("/messages", handleRequests)
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rest.GetMessage(w, r)
	case "POST":
		rest.CreateMessage(w, r)
	case "PUT":
		rest.UpdateMessage(w, r)
	case "DELETE":
		rest.DeleteMessage(w, r)
	default:
		http.Error(w, "Unsupported request method.", http.StatusMethodNotAllowed)
	}
}
