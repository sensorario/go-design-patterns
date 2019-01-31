package main

import "fmt"
import "os"

type TrafficLightState interface {
	Exec(k *AContext) bool
	Name() string
}

type FinishState struct{}

func (s *FinishState) Exec(k *AContext) bool {
	fmt.Println("FINISHED !!!")
	return false
}
func (s *FinishState) Name() string {
	return "finish"
}

type EndState struct{}

func (s *EndState) Exec(k *AContext) bool {
	if k.Exit == true {
		k.CurrentState = &FinishState{}
		return true
	}

	k.CurrentState = &Ask{}
	return true
}

func (s *EndState) Name() string {
	return "end"
}

type AContext struct {
	CurrentState TrafficLightState
	Number       int
	Exit         bool
}

func (a *AContext) prntState() {
	fmt.Println(">>", a.CurrentState.Name())
	if a.CurrentState.Name() == "end" {
		fmt.Println("")
	}
}

type Ask struct{}

func (s *Ask) Exec(k *AContext) bool {
	var n int
	fmt.Print(">> ")
	fmt.Fscanf(os.Stdin, "%v", &n)
	if n == k.Number {
		k.Exit = true
	} else {
		if n > k.Number {
			fmt.Println(">> you number is greater")
		} else {
			fmt.Println(">> you number is lower")
		}
	}
	k.CurrentState = &EndState{}
	return true
}

func (s *Ask) Name() string {
	return "guess the number ... "
}
