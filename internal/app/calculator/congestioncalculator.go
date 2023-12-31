package calculator

import (
	"time"

	"github.com/gustavocmaciel/go-road-fare/internal/app/vehicles"
)

func GetTax(vehicle vehicles.Vehicle, dates []time.Time) int {
	intervalStart := dates[0]
	totalFee := 0
	tempFee := 0
	for _, date := range dates {
		nextFee := getTollFee(date, vehicle)

		diffInNanos := date.UnixNano() - intervalStart.UnixNano()
		minutes := diffInNanos / 1000000 / 1000 / 60

		if minutes <= 60 {
			if totalFee > 0 {
				totalFee = totalFee - tempFee
			}
			if nextFee >= tempFee {
				tempFee = nextFee
			}
		} else {
			totalFee = totalFee + nextFee
		}
	}

	if totalFee > 60 {
		totalFee = 60
	}
	return totalFee
}

func isTollFreeVehicle(v vehicles.Vehicle) bool {
	if v == nil {
		return false
	}
	vehicleType := v.GetVehicleType()

	return vehicleType == vehicles.TollFreeVehicles(vehicles.Motorcycle).String() || vehicleType == vehicles.TollFreeVehicles(vehicles.Tractor).String() || vehicleType == vehicles.TollFreeVehicles(vehicles.Emergency).String() || vehicleType == vehicles.TollFreeVehicles(vehicles.Diplomat).String() || vehicleType == vehicles.TollFreeVehicles(vehicles.Foreign).String() || vehicleType == vehicles.TollFreeVehicles(vehicles.Military).String()
}

func getTollFee(t time.Time, v vehicles.Vehicle) int {
	if isTollFreeDate(t) || isTollFreeVehicle(v) {
		return 0
	}

	hour, minute := t.Hour(), t.Minute()

	if hour == 6 && minute >= 0 && minute <= 29 {
		return 8
	}
	if hour == 6 && minute >= 30 && minute <= 59 {
		return 13
	}
	if hour == 7 && minute >= 0 && minute <= 59 {
		return 18
	}
	if hour == 8 && minute >= 0 && minute <= 29 {
		return 13
	}
	if hour >= 8 && hour <= 14 && minute >= 30 && minute <= 59 {
		return 8
	}
	if hour == 15 && minute >= 0 && minute <= 29 {
		return 13
	}
	if hour == 15 && minute >= 0 || hour == 16 && minute <= 59 {
		return 18
	}
	if hour == 17 && minute >= 0 && minute <= 59 {
		return 13
	}
	if hour == 18 && minute >= 0 && minute <= 29 {
		return 8
	}

	return 0
}

func isTollFreeDate(date time.Time) bool {
	year := date.Year()
	month := date.Month()
	day := date.Day()

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}

	if year == 2013 {
		if month == 1 && day == 1 || month == 3 && (day == 28 || day == 29) || month == 4 && (day == 1 || day == 30) || month == 5 && (day == 1 || day == 8 || day == 9) || month == 6 && (day == 5 || day == 6 || day == 21) || month == 7 || month == 11 && day == 1 || month == 12 && (day == 24 || day == 25 || day == 26 || day == 31) {
			return true
		}
	}
	return false
}
