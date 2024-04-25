// File: rest/rest.go

package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Message struct to encode the message as JSON
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Path string `json:"path,omitempty"`
}

var (
	messages map[int]*Message
	mu       sync.Mutex
	nextID   int
)

func init() {
	messages = make(map[int]*Message)
	nextID = 1
}

// GetMessage retrieves a message by ID
func GetMessage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requested URL: %s", r.URL.Path)
	idStr := r.URL.Path[len("/messages/"):]
	log.Printf("Extracted ID: %s", idStr)

	id, err := strconv.Atoi(r.URL.Path[len("/messages/"):])
	if err != nil {
		http.Error(w, "Invalid message ID maybe baby", http.StatusBadRequest)
		return
	}

	mu.Lock()

	if messages == nil {
		fmt.Fprintf(w, "messages is nil")
	}

	message, ok := messages[id]
	mu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(message)
}

// CreateMessage creates a new message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	msg.ID = nextID
	messages[nextID] = &msg
	nextID++
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

// UpdateMessage updates an existing message
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/messages/"):])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}
	var updatedMsg Message
	if err := json.NewDecoder(r.Body).Decode(&updatedMsg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	msg, ok := messages[id]
	if !ok {
		mu.Unlock()
		http.NotFound(w, r)
		return
	}
	updatedMsg.ID = id
	*msg = updatedMsg
	mu.Unlock()
	json.NewEncoder(w).Encode(updatedMsg)
}

// DeleteMessage deletes a message
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/messages/"):])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}
	mu.Lock()
	_, ok := messages[id]
	if !ok {
		mu.Unlock()
		http.NotFound(w, r)
		return
	}
	delete(messages, id)
	mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}
