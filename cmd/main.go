package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocmaciel/go-road-fare/internal/app/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Calculator")
	})
	router.HandleFunc("/api/calculator/{vehicle}/{time}", handlers.GetVehicleTax).Methods(http.MethodGet)
	log.Println("API is running!")
	http.ListenAndServe("localhost:8080", router)
	// http.ListenAndServe(":8080", router) // for production environment

}
