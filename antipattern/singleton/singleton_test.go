package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	firstInstance := GetInstance()

	if firstInstance.NumberOfCreations() != 1 {
		t.Error("expected just one number of creations")
	}

	secondInstance := GetInstance()
	if firstInstance != secondInstance {
		t.Error("expected same instance")
	}

	thirdInstance := GetInstance()
	if thirdInstance.NumberOfCalls() != 3 {
		t.Error("expected three calls")
	}

	if firstInstance.NumberOfCreations() != 1 {
		t.Error("expected just one number of creations")
	}
}
