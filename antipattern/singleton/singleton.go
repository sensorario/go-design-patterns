package singleton

type singleton struct {
	calls int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	instance.calls++
	return instance
}

func (s *singleton) NumberOfCalls() int {
	return s.calls
}
