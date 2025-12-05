package main

import (
	"net"
	"io"
	"net/http"
	"os"
	"log"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	var message string = "Hello, World!\n";
	log.Println("root requested. Response: ", message)
	io.WriteString(w, message)
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Health check requested")
}


func loadPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
        log.Println("Using default port 8080")
    }
    return port
}

func main() {
	var port string = loadPort();
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/healthz", getHealth)
	address := net.JoinHostPort("", port)
	log.Println("Server is running on port", port)
	http.ListenAndServe(address, nil)
}