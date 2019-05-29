# Flyweight

## Elements

*to do*

## Dscription

This pattern it's very commonly used in computer graphics and the video game
industry. It allow sharing the state of a heavy object between many instances of
some type.

This example uses a creational pattern to create objects instance.

```go
func TestFactoryCreatesObjects(t *testing.T) {
	f := NewObjectFactory()
	firstObject := f.GetObject(TYPE_ONE)
	if firstObject == nil {
		t.Error("The pointer to the TYPE_ONE was nil")
	}
}

```

According to the flyweight pattern, each object requested is returned. Well,
what is really returned is not an object but a pointer to that object.

```go
func TestFactoryCreatesTwoObjects(t *testing.T) {
	f := NewObjectFactory()
	_ = f.GetObject(TYPE_ONE)
	secondObject := f.GetObject(TYPE_ONE)
	if secondObject == nil {
		t.Error("The pointer to the TYPE_ONE was nil")
	}
}

```

Each pointer created is a different pointer.

```go
func TestFactoryCreatesJustObjectOfTypes(t *testing.T) {
	f := NewObjectFactory()
	firstObject := f.GetObject(TYPE_ONE)
	secondObject := f.GetObject(TYPE_ONE)
	if firstObject != secondObject {
		t.Error("TYPE_ONE pointers weren't the same")
	}
}

```

Even if object of TYPE_ONE is requested more times, the number of created
objects is equals to the number of type requested.

```go
func TestNumberOfObjectsIsAlwaysNumberOfTypeOfObjectCreated(t *testing.T) {
	f := NewObjectFactory()
	_ = f.GetObject(TYPE_ONE)
	_ = f.GetObject(TYPE_ONE)
	if f.GetNumberOfObjects() != 1 {
		t.Errorf(
			"The number of objects created was not 1: %d\n",
			f.GetNumberOfObjects(),
		)
	}
}

```

Finally, and for completeness, if two objects are requested a huge amount of
time, the number of object created is still two.

```go
func TestHighVolume(t *testing.T) {
	f := NewObjectFactory()
	objects := make([]*Object, 500000*2)
	for i := 0; i < 500000; i++ {
		objects[i] = f.GetObject(TYPE_ONE)
	}
	for i := 500000; i < 2*500000; i++ {
		objects[i] = f.GetObject(TYPE_TWO)
	}
	if f.GetNumberOfObjects() != 2 {
		t.Errorf(
			"The number of objects created was not 2: %d\n",
			f.GetNumberOfObjects(),
		)
	}
}
```
