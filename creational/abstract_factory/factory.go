package factory

import (
	"fmt"
)

const (
	ItalianType = 1
)

const (
	FerrariModel          = 1
	CarWithFiveWheelModel = 2
)

type Vehicle interface {
	NumOfWheels() int
	GetModelName() string
}

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

type ItalianFactory struct{}

type CarWithFiveWheelType struct{}

func (f *CarWithFiveWheelType) NumOfWheels() int {
	return 5
}

func (f *CarWithFiveWheelType) GetModelName() string {
	return "Star"
}

type FerrariModelType struct {
}

func (f *FerrariModelType) NumOfWheels() int {
	return 4
}

func (f *FerrariModelType) GetModelName() string {
	return "Ferrari"
}

func (i *ItalianFactory) Build(v int) (Vehicle, error) {
	switch v {
	case FerrariModel:
		return new(FerrariModelType), nil
	case CarWithFiveWheelModel:
		return new(CarWithFiveWheelType), nil
	}
	return nil, fmt.Errorf("No Italian cars of type %d\n", v)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case ItalianType:
		return new(ItalianFactory), nil
	default:
		return nil, fmt.Errorf("No factory with id %d\n", f)
	}
}
