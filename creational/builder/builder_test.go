package builder

import "testing"
import "strconv"

func TestBuilderPattern(t *testing.T) {
	director := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	director.SetBuilder(carBuilder)
	director.Construct()
	car := carBuilder.Build()
	if car.Wheels != 4 {
		t.Errorf("Something went wrong : " + strconv.Itoa(car.Wheels) + " wheels found")
	}

	biciclettaBuilder := &BiciclettaBuilder{}
	director.SetBuilder(biciclettaBuilder)
	director.Construct()
	bike := biciclettaBuilder.Build()
	if bike.Wheels != 2 {
		t.Errorf("Something went wrong : " + strconv.Itoa(bike.Wheels) + " wheels found")
	}
}
