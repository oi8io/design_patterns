package creation

import "testing"

func TestVehicleDirector_BuildCar(t *testing.T) {
	dor := new(VehicleDirector)
	b := &CarBuilder{}
	dor.Build(b)
}
func TestVehicleDirector_BuildBike(t *testing.T) {
	dor := new(VehicleDirector)
	b := &BikeBuilder{}
	dor.Build(b)
}
