# Creational » Pool

## Description

The object pool design pattern creates a set of objects that may be reused. When a new object is needed, it is requested from the pool. Once an object has been used and returned, existing references will become invalid. If a previously prepared object is available it is returned immediately, avoiding the instantiation cost. If no objects are present in the pool, a new item is created and returned.

## Implementation

For this pattern we will use two object. The PoolObject and the Pool. The former is the object that must be managed, the latter is the object that manage the pool. That provide idle objects. That receive object no more in use.

```go
type PoolObject struct {
	id int
}

type Pool struct {
	idle   *list.List
	active *list.List
}
```

In this implementation, the pool initialize two lists. The idle list will contains all retured object, read for reuse. The active list, contains all generate object.

```go
func InitPool() Pool {
	pool := &Pool{
		list.New(),
		list.New(),
	}
	return *pool
}
```

In a very simple implementation, object are created each time requested from the pool.

```go
func (p *Pool) Loan() PoolObject {
	object := PoolObject{p.NumberOfObjectsInPool() + 1}
	p.active.PushBack(object)
	return object
}
```

The complete implementation of Loan method checl if there are already created object. If idle object exists, it is provided avoiding creation time.

```go
func (p *Pool) Loan() PoolObject {
	if p.idle.Len() > 0 {
		for e, i := p.idle.Front(), 0; e != nil; e, i = e.Next(), i+1 {
			if i == 0 {
				object := e.Value.(PoolObject)
				return object
			}
		}
	}

	object := PoolObject{p.NumberOfObjectsInPool() + 1}
	p.active.PushBack(object)
	return object
}
```

When an object is used, is returned to the pool. Used object are stored in idle list, ready to be returned.

```go
func (p *Pool) Receive(object PoolObject) {
	p.idle.PushBack(object)
	for e, i := p.active.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		if object == e.Value.(PoolObject) {
			p.active.Remove(e)
			return
		}
	}
}
```

## Usage

First, the pool provide two object. Second, the pool receive one used object. Third, another loan is requested but instead of provide another instance of new object, thirdObject contains firstObject. No time is spent to build another object.

```go
pool := InitPool()
firstObject := pool.Loan()
secondObject := pool.Loan()

pool.Receive(firstObject)
thirdObject := pool.Loan()

if firstObject.id != thirdObject.id {
  panic("thir object must contain firs one")
}
```
