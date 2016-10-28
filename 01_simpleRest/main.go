package main

import (
	"log"
	"net/http"

	"app/handlers"
	"app/storage"
)

func main() {

	db := storage.NewInMemoryDatabase()
	mux := http.NewServeMux()

	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.SetKey(db))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
