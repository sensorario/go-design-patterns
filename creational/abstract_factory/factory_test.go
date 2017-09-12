package factory

import (
	"testing"
)

func TestFoo(t *testing.T) {
	itaF, _ := BuildFactory(ItalianType)
	m, _ := itaF.Build(FerrariModel)

	car, ok := m.(Vehicle)
	if !ok {
		t.Fatal("Invalid model")
	}

	t.Logf("%v car has %d wheels", car.GetModelName(), car.NumOfWheels())
}

func TestBar(t *testing.T) {
	itaF, _ := BuildFactory(ItalianType)
	m, _ := itaF.Build(CarWithFiveWheelModel)

	car, ok := m.(Vehicle)
	if !ok {
		t.Fatal("Invalid model")
	}

	if m.NumOfWheels() != 5 {
		t.Fatal("Star car should have 5 wheels")
	}

	t.Logf("%v car has %d wheels", car.GetModelName(), car.NumOfWheels())
}
