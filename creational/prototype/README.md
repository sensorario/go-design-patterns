# Creational » Prototype

## Intent

 - provide an interface for creating families of related or dependent objects
 - a hierarchy that encapsulates: many possible "platforms", and the construction of a suite of "products".
 - the new operator considered harmful.

## Description

With this pattern, is used an already created instance of some type to clone it and complete it with the particular needs of each context. Objects to clone are created at compilation time and can be cloned as many times it is needed at runtime.

```go
firstInstance := GetInstance()

secondInstance := GetInstance()
if firstInstance != secondInstance {
	t.Error("expected same instance")
}
```

Any following calls, update first intance.

```go
thirdInstance := GetInstance()
if thirdInstance.NumberOfCalls() != 3 {
	t.Error("expected three calls")
}
```

## Few words …

Is not good pattern if used to bring the state of the applicatoin and can cnange during its lifecycle. Making something global to avoid passing it around is a code smell. But use it to read configuration is good. Used to load a resource just first time is requested and to provide that resource everywere is a good way to use this pattern.

## Usage

Two kind of t-shirt are needed. Same model, with different SKU. Instead of recreeate same white shirt from scratch, base model is cloned.

```go
shirtCache := GetShirtsCloner()
firstItem, err := shirtCache.GetClone(White)
firstItem.SKU = "abc"

secondItem, err := shirtCache.GetClone(White)
secondItem.SKU = "xxx"
```
