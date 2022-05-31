package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("APP:", os.Getenv("APP"))
		appName := os.Getenv("APP")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Inicio"+appName)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("APP:", os.Getenv("APP"))
		appName := os.Getenv("APP")
		w.WriteHeader(http.StatusOK)
		userName := r.URL.Query().Get("name")
		if userName != "" {
			io.WriteString(w, appName+":Username is"+userName)
		}
	})

	err := http.ListenAndServe(":8081", nil)

	log.Print("Server is listening on port :8081")
	if err != nil {
		log.Fatalf("Could not start the server: %s\n", err.Error())
	}
}
