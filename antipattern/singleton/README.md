# Creational » Singleton

## Intent

 - guarantee that a class has only one instance
 - provide access to that instance

## Description

The Singleton pattern aims to provide same instance of an objefct and guarantee that there are no duplicates. At the first call the instance is created. Following calls return first instance.

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
