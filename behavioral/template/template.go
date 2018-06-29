package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type TheTemplate interface {
	first() string
	second() string
	customStep(MessageRetriever) string
}

type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) second() string {
	return "template"
}

func (t *Template) customStep(m MessageRetriever) string {
	return strings.Join(
		[]string{
			t.first(),
			m.Message(),
			t.second(),
		},
		" ",
	)
}

type Anonymous struct{}

func (a *Anonymous) first() string {
	return "hello"
}

func (a *Anonymous) second() string {
	return "template"
}

func (a *Anonymous) customStep(f func() string) string {
	return strings.Join(
		[]string{
			a.first(),
			f(),
			a.second(),
		},
		" ",
	)
}

type Wrapper struct {
	myFunc func() string
}

func (a *Wrapper) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &Wrapper{myFunc: f}
}
