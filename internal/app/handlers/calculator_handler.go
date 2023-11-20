package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gustavocmaciel/go-road-fare/internal/app/calculator"
	"github.com/gustavocmaciel/go-road-fare/internal/app/vehicles"
)

type GetTaxResponse struct {
	TaxAmount int `json:"taxAmount"`
}

func GetVehicleTax(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleArg := vars["vehicle"]
	timeArg := vars["time"]

	var vehicle vehicles.Vehicle

	timeArray := strings.Split(timeArg, ",")

	var dates []time.Time

	for _, t := range timeArray {

		parsedTime, err := time.Parse("2006-01-02T15:04:05", t)
		if err != nil {
			errorMessage := "Error parsing time string"
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}
		dates = append(dates, parsedTime)
	}

	switch vehicleArg {
	case "car":
		vehicle = vehicles.Car{}
	case "motorbike":
		vehicle = vehicles.Motorbike{}
	case "motorcycle":
		vehicle = vehicles.TollFreeVehicleType{TollFreeVehicles: vehicles.Motorcycle}
	case "tractor":
		vehicle = vehicles.TollFreeVehicleType{TollFreeVehicles: vehicles.Tractor}
	case "emergency":
		vehicle = vehicles.TollFreeVehicleType{TollFreeVehicles: vehicles.Emergency}
	case "foreign":
		vehicle = vehicles.TollFreeVehicleType{TollFreeVehicles: vehicles.Foreign}
	case "military":
		vehicle = vehicles.TollFreeVehicleType{TollFreeVehicles: vehicles.Military}
	default:
		vehicle = nil
	}

	tax := calculator.GetTax(vehicle, dates)
	response := GetTaxResponse{tax}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
