package singleton

type singleton struct {
	calls     int
	creations int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
		instance.creations = 1
	}
	instance.calls++
	return instance
}

func (s *singleton) NumberOfCalls() int {
	return s.calls
}

func (s *singleton) NumberOfCreations() int {
	return s.creations
}
