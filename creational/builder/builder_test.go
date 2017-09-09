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

	bicicletta := &BiciclettaBuilder{}
	director.SetBuilder(bicicletta)
	director.Construct()
	biciletta := bicicletta.Build()
	if biciletta.Wheels != 2 {
		t.Errorf("Something went wrong : " + strconv.Itoa(bicicletta.Wheels) + " wheels found")
	}
}
