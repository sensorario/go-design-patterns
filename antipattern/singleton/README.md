# Creational » Singleton

## Description

The Singleton pattern aims to provide same instance of an object and guarantee
that there are no duplicates. At the first call the instance is created.
Following calls return first instance.

```go
firstInstance := GetInstance()

secondInstance := GetInstance()
if firstInstance != secondInstance {
	t.Error("expected same instance")
}
```

Any following calls, update first instance.

```go
thirdInstance := GetInstance()
if thirdInstance.NumberOfCalls() != 3 {
	t.Error("expected three calls")
}
```

## Few words …

Is not good pattern if used to bring the state of the application and can
change during its life cycle. Making something global to avoid passing it
around is a code smell. But use it to read configuration is good. Used to load
a resource just first time is requested and to provide that resource every were
is a good way to use this pattern.
