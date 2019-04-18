package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 || car.Seats != 5 || car.Structure != "Car" {
		t.Error("got errors: car builder crashed")
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 || bike.Seats != 2 || bike.Structure != "Motorbike" {
		t.Error("go errors: bike builder crashed")
	}
}
