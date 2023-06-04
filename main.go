package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	// simply a file server
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "80"
	}
	fmt.Println("listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
