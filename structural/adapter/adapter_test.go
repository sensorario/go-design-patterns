package main

import "testing"

func TestAdater(t *testing.T) {
	msg := "Hello world!!!"
	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello world!!!\n" {
		t.Errorf("Message didnt match: %s\n", returnedMsg)
	}
}
