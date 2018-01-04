package flyweight

import (
	"testing"
)

func TestFactoryCreatesObjects(t *testing.T) {
	f := NewObjectFactory()
	firstObject := f.GetObject(TYPE_ONE)
	if firstObject == nil {
		t.Error("The pointer to the TYPE_ONE was nil")
	}
}

func TestFactoryCreatesTwoObjects(t *testing.T) {
	f := NewObjectFactory()
	_ = f.GetObject(TYPE_ONE)
	secondObject := f.GetObject(TYPE_ONE)
	if secondObject == nil {
		t.Error("The pointer to the TYPE_ONE was nil")
	}
}

func TestFactoryCreatesJustObjectOfTypes(t *testing.T) {
	f := NewObjectFactory()
	firstObject := f.GetObject(TYPE_ONE)
	secondObject := f.GetObject(TYPE_ONE)
	if firstObject != secondObject {
		t.Error("TYPE_ONE pointers weren't the same")
	}
}

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
