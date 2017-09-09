package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	firstInstance := GetInstance()

	secondInstance := GetInstance()
	if firstInstance != secondInstance {
		t.Error("expected same instance")
	}

	thirdInstance := GetInstance()
	if thirdInstance.NumberOfCalls() != 3 {
		t.Error("expected three calls")
	}
}
