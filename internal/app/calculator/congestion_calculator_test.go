package calculator

import (
	"testing"
	"time"

	"github.com/gustavocmaciel/go-road-fare/internal/app/vehicles"
)

func TestGetTax(t *testing.T) {
	// Mock data
	vehicle := vehicles.Car{} // Replace with the actual type of vehicle
	dates := []time.Time{
		parseTime("2023-05-01T10:05:00"),
		parseTime("2023-05-01T10:20:00"),
		parseTime("2023-05-01T17:40:00"),
		parseTime("2023-05-01T18:00:00"),
	}

	// Call the function
	result := GetTax(vehicle, dates)

	// Expected result
	expected := 21 // Replace with the expected result based on your test case

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Helper function to parse time string
func parseTime(timeStr string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05", timeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func TestIsTollFreeVehicle(t *testing.T) {
	// Test with a toll-free vehicle (Motorcycle)
	tollFreeMotorcycle := vehicles.TollFreeVehicleType{TollFreeVehicles: vehicles.Motorcycle}
	resultTollFreeMotorcycle := isTollFreeVehicle(tollFreeMotorcycle)
	if !resultTollFreeMotorcycle {
		t.Errorf("Expected true for toll-free Motorcycle, but got false")
	}

	// Test with a non-toll-free vehicle (Car)
	nonTollFreeCar := vehicles.Car{}
	resultNonTollFreeCar := isTollFreeVehicle(nonTollFreeCar)
	if resultNonTollFreeCar {
		t.Errorf("Expected false for non-toll-free Car, but got true")
	}

	// Test with a nil vehicle
	resultNilVehicle := isTollFreeVehicle(nil)
	if resultNilVehicle {
		t.Errorf("Expected false for nil vehicle, but got true")
	}

}
