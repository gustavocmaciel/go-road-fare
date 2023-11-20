package vehicles

import "fmt"

type TollFreeVehicles int

const (
	Motorcycle TollFreeVehicles = iota
	Tractor    TollFreeVehicles = iota
	Emergency  TollFreeVehicles = iota
	Diplomat   TollFreeVehicles = iota
	Foreign    TollFreeVehicles = iota
	Military   TollFreeVehicles = iota
)

func (tfv TollFreeVehicles) String() string {
	switch tfv {
	case Motorcycle:
		return "Motorcycle"
	case Tractor:
		return "Tractor"
	case Emergency:
		return "Emergency"
	case Foreign:
		return "Foreign"
	case Military:
		return "Military"
	default:
		return fmt.Sprintf("%d", int(tfv))
	}
}

type TollFreeVehicleType struct {
	TollFreeVehicles
}

func (c TollFreeVehicleType) GetVehicleType() string {
	return c.TollFreeVehicles.String()
}
