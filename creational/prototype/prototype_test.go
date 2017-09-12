package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	firstItem, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	if firstItem == whitePrototype {
		t.Fatal("firstitem cannot be equal to the white prototype")
	}

	secondItem, err := shirtCache.GetClone(White)

	firstShirt, ok := firstItem.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldnt be done successfully")
	}

	firstShirt.SKU = "abbcde"
	secondShirt, ok := secondItem.(*Shirt)

	if firstShirt.SKU == secondShirt.SKU {
		t.Fatal("SKU's must be different")
	}
}
