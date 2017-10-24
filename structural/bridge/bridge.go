package main

import (
	"errors"
	"fmt"
	"io"
)

type PrinterImp1 struct{}

func (p *PrinterImp1) PrintMessage(msg string) error {
	fmt.Println("%s\n", msg)
	return nil
}

// ---

type PrinterImp2 struct {
	Writer io.Writer
}

func (d *PrinterImp2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImp2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

// ---

type PrinterAPI interface {
	PrintMessage(string) error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}
