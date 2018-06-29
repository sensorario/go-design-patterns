# Template

## Description

Strategy pattern consists in a different way to solve same problem with
different strategies. In Strategy entire algorithm is store in different
classes. In Template pattern, instead, the algorithm is just one. The template
defines steps but some other steps are deferred to user.

## Implementation

In this case the template is a sort of flow of steps to solve a problem. While
in strategy pattern each strategy implements entire solution, in this case the
template define the steps execution but defer one or more steps to the user. In
this case, the deferred step is called `templateSteps`.

```go
type TheTemplate interface {
	first() string
	second() string
	templateSteps(MessageRetriever) string
}
```

The `MessageRetriever` is an interface with `Message()` method. The template
pattern here will joins strings. While first and third steps are implemented,
the custom step is implemented by user.

```go
type MessageRetriever interface {
	Message() string
}
```

Now let's implement a concrete `Template`.

```go
type Template struct{}
```

In this case first and second steps are defined. The `templateSteps` will join
different strings receiving the third step from outside.

```go
func (t *Template) first() string {
	return "hello"
}

func (t *Template) second() string {
	return "template"
}
```

```go
func (t *Template) templateSteps(m MessageRetriever) string {
	return strings.Join(
		[]string{
			t.first(),
			m.Message(),
			t.second(),
		},
		" ",
	)
}
```
