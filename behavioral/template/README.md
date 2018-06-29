# Template

## Description

Strategy pattern consists in different way to solve same problem. In Strategy
entire algorithm is store in different classes. In Template pattern the
algorithm is just one. The template define steps but some steps are deferred to
user.

This kind of pattern is widely used in MVC web frameworks when steps are always
the same:

 - Handle Request
 - Authentication
 - Authorization
 - Retrieve data from DB
 - Do something
 - Send Response

All these steps are execute for each http request. Just the `do something`
part is defined by the user.

## Implementation

```go
type TranslatorRing interface {
	Next(TranslatorRing)
	GetNext() TranslatorRing
	KnowsWord(s string) bool
	TranslationOf(s string) string
}
```
