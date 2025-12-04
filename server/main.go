package main

import (
	"net"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"log"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	var message string = "Hello, World!\n";
	log.Println("root requested. Response: ", message)
	io.WriteString(w, message)
}

func loadPortFromEnvFile() string {
		err := godotenv.Load()
		port := ""
    if err != nil {
    	log.Println("Error loading .env file: ", err)
			port = "8080"
			log.Println("Using default port 8080")
	  } else {
			port = os.Getenv("PORT")
		}
	return port
}

func main() {
	var port string = loadPortFromEnvFile();
	http.HandleFunc("/", getRoot)
	address := net.JoinHostPort("", port)
	log.Println("Server is running on port", port)
	http.ListenAndServe(address, nil)
}