package builder

type Product struct {
	Wheels int
	Seats  int
}

type VehicleBuildProcess interface {
	SetWheelsNumber() VehicleBuildProcess
	SetSeatsNumber() VehicleBuildProcess
	GetVehicle() Product
}

type ManufacturingDirector struct {
	builder VehicleBuildProcess
}

func (f *ManufacturingDirector) SetBuilder(b VehicleBuildProcess) {
	f.builder = b
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeatsNumber().SetWheelsNumber()
}

type CarBuilder struct {
	p Product
}

func (c *CarBuilder) SetWheelsNumber() VehicleBuildProcess {
	c.p.Seats = 5
	return c
}

func (c *CarBuilder) SetSeatsNumber() VehicleBuildProcess {
	c.p.Wheels = 4
	return c
}

func (c *CarBuilder) GetVehicle() Product {
	return c.p
}

func (c *CarBuilder) Build() Product {
	return c.p
}

type BiciclettaBuilder struct {
	p Product
}

func (c *BiciclettaBuilder) SetWheelsNumber() VehicleBuildProcess {
	c.p.Seats = 1
	return c
}

func (c *BiciclettaBuilder) SetSeatsNumber() VehicleBuildProcess {
	c.p.Wheels = 2
	return c
}

func (c *BiciclettaBuilder) GetVehicle() Product {
	return c.p
}

func (c *BiciclettaBuilder) Build() Product {
	return c.p
}
