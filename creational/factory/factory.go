package factory

import "errors"

const (
	Italian = 1
	English = 2
)

type Translator interface {
	GetGreeting() string
}

func GetTranslator(m int) (Translator, error) {
	switch m {
	case Italian:
		return new(ItalianGreeting), nil
	case English:
		return new(EnglishGreeting), nil
	default:
		return nil, errors.New("Unknown building")
	}
	return nil, errors.New("Not implemented yet")
}

type ItalianGreeting struct{}

func (c *ItalianGreeting) GetGreeting() string {
	return "Ciao"
}

type EnglishGreeting struct{}

func (d *EnglishGreeting) GetGreeting() string {
	return "Hello"
}
