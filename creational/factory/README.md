# Creational » Factory

## Description

The purpose of current pattern is to provide an interface that fits the developer's needs, delegating the decision of objects creation to a factory. Only the factory knows how to create stuffs. The creation project is completely abstract here.

## Implementation

In this example we will build a sort of greeting generators. Just for italian and english translations.

First of all we need languages. We can define them as constants:

```go
const (
	Italian = 1
	English = 2
)
```

Then, we need a translator. A translator that build a greeting message.

```go
type Translator interface {
	GetGreeting() string
}
```

To be sure that all works fine, we create a greeting from Italian. In our test we just care about the content of the message.

```go
func TestCreateGetGreetingOfWarehouse(t *testing.T) {
	greeting, err := GetTranslator(Italian)
	if err != nil {
		t.Fatal("A GetGreetingment method of type 'Italian' must exists")
	}

	msg := greeting.GetGreeting()
	if !strings.Contains(msg, "Ciao") {
		t.Error("The italian greeting isn't correct")
	}
}
```

As you can see, from tests EnglishGreeting and ItalianGreeting are not present. The factory successfully abstract the message creation, delegating it to ItalianGreeting and EnglishGreeting.
