package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(BesicHndler),
	}

	log.Print("servser is runnign")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to listen to server")
	}
}

func BesicHndler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
