package main

import (
	"io"
	"log"
	"math/rand/v2"
	"net"
	"net/http"
	"os"
	"strconv"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	var message string = "Hello, World!\n"
	log.Println("root requested. Response: ", message)
	io.WriteString(w, message)
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Health check requested")
}

func getReadiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Readiness check requested")
}

func getRandomNumber(w http.ResponseWriter, r *http.Request) {
	var number int = rand.IntN(100) + 500
	log.Println("Random number requested")
	io.WriteString(w, strconv.Itoa(number))

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
	var port string = loadPort()
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/healthz", getHealth)
	http.HandleFunc("/readyz", getReadiness)
	// http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/metrics", getRandomNumber)
	address := net.JoinHostPort("", port)
	log.Println("Server is running on port", port)
	http.ListenAndServe(address, nil)
}
