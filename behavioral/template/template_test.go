package template

import (
	"strings"
	"testing"
)

type EmbedTemplate struct {
	Template
}

func (m *EmbedTemplate) Message() string {
	return "world"
}

func TestCustomStepIsCalled(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &EmbedTemplate{}
		res := s.customStep(s)
		check(res, " world ", t)
	})

	t.Run("Define custom step via anonymous function", func(t *testing.T) {
		m := new(Anonymous)
		res := m.customStep(func() string {
			return "world"
		})
		check(res, " world ", t)
	})

	t.Run("Using anonymous functions adapted to an interface", func(t *testing.T) {
		customStepStep := MessageRetrieverAdapter(func() string {
			return "world"
		})
		if customStepStep == nil {
			t.Fatal("Can not continue with a nil custom step")
		}
		template := Template{}
		res := template.customStep(customStepStep)
		check(res, " world ", t)
	})
}

func check(res string, expected string, t *testing.T) {
	if !strings.Contains(res, expected) {
		t.Errorf("Expected string '%s' was not found on returned string '%s'\n", expected, res)
	}
}
