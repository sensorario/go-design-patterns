package main

import "fmt"
import "sync"
import "errors"
import "math/rand"
import "time"

type SuccessFunc func(string)
type FailureFunc func(error)
type ExecuteFunc func(int) (string, error)

type Subject struct {
	success SuccessFunc
	failure FailureFunc
}

func (s *Subject) Success(f SuccessFunc) *Subject {
	s.success = f
	return s
}

func (s *Subject) Failure(f FailureFunc) *Subject {
	s.failure = f
	return s
}

func (s *Subject) Execute(f ExecuteFunc) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func(s *Subject) {
		r := rand.NewSource(time.Now().UnixNano())
		n := rand.New(r)

		str, err := f(n.Intn(200))

		if err != nil {
			s.failure(err)
			wg.Done()
		} else {
			s.success(str)
			wg.Done()
		}
	}(s)

	wg.Wait()
}

func main() {
	s := Subject{}

	s.Success(func(m string) {
		fmt.Println("SUCCESS: ", m)
	}).Failure(func(e error) {
		fmt.Println("FAILURE: ", e)
	}).Execute(func(num int) (string, error) {
		if num < 100 {
			return "", errors.New("too low number")
		} else {
			return "valid number", nil
		}
	})

}
