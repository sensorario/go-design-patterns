# Creational » Builder

## Intent

 - separate the construction of a complex object from its representation;
 - parse a complex representation, create one of several targets;

## Description

The Builder is a pattern used to build objects. Objects in go can be created with just {}. But in Go is possible to create objects composed by other objects. This is really idiomatic in go, as it doesn't support inheritance.

The target of this pattern is to build something. An in this example we will create a Product.

```go
type Product struct {
	Wheels int
	Seats  int
}
```

This kind of product have wheels and seats. All products must be builded following a process. This process aims to complete the product. The actions to do to complete the product are `SetWheelsNumber` and `SetSeatsNumber`. After this, the process can deploy the product.

```go
type VehicleBuildProcess interface {
	SetWheelsNumber() VehicleBuildProcess
	SetSeatsNumber() VehicleBuildProcess
	GetVehicle() Product
}
```

All the process are managed by a ManufacturingDirector. It will ask the builder all the actions to do.

```go
type ManufacturingDirector struct {
	builder VehicleBuildProcess
}
```

In the following code we tell to the director "who" is the builder and he will make the product with construct method.

```go
func (f *ManufacturingDirector) SetBuilder(b VehicleBuildProcess) {
	f.builder = b
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeatsNumber().SetWheelsNumber()
}
```

## A car builder

The final thing to do is to create a builder. In this example we will build a car. A car with 5 seats and 4 wheels.

```go
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
```

# Add new product, …

After we have built a car, we could create a bicycle. Here the test:

```go
bicicletta := &BiciclettaBuilder{}
director.SetBuilder(bicicletta)
director.Construct()
biciletta := bicicletta.Build()
if biciletta.Wheels != 2 {
  t.Errorf("Something went wrong : " + strconv.Itoa(bicicletta.Wheels) + " wheels found")
}
```

And here the builder:

```go
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
```

## Usage

The director must build a vehicle like a car. Once the builder is provided. The director ask the builder to build the car. Finally we'll have right model of car.

```go
director := ManufacturingDirector{}

carBuilder := &CarBuilder{}
director.SetBuilder(carBuilder)
director.Construct()
car := carBuilder.Build()
```

The director must build a vehicle like a bike. Once the builder is provided. The director ask the builder to build the bike. Finally we'll have right model of bike.

```go
biciclettaBuilder := &BiciclettaBuilder{}
director.SetBuilder(biciclettaBuilder)
director.Construct()
bike := biciclettaBuilder.Build()
```
