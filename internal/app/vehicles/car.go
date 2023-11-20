package vehicles

type Car struct {
	Type string
}

func (c Car) GetVehicleType() string {
	return "Car"
}
