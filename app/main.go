package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Inicio\n")
	})

	err := http.ListenAndServe(":8081", nil)

	log.Print("Server is listening on port :8081")
	if err != nil {
		log.Fatalf("Could not start the server: %s\n", err.Error())
	}
}
