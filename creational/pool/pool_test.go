package pool

import (
	"strconv"
	"strings"
	"testing"
)

func TestWhenObjectIsNeededItIsRequestedFromThePool(t *testing.T) {
	pool := InitPool()
	if pool.NumberOfActiveObjects() != 0 {
		t.Fatal("Actually there should be zero objects in pool")
	}
	if pool.NumberOfIdleObjects() != 0 {
		t.Fatal("Actually there should be zero idle objects in pool")
	}
	if pool.NumberOfObjectsInPool() != 0 {
		t.Fatal("There shouldnt be more than zero object in pool")
	}

	objectOne := pool.Loan()
	if pool.NumberOfActiveObjects() != 1 {
		t.Fatal("Something went wrong")
	}
	if pool.NumberOfIdleObjects() != 0 {
		t.Fatal("Actually there should be zero idle objects in pool")
	}

	createdId := objectOne.id
	if createdId != 1 {
		t.Fatal("object id should be 1")
	}

	pool.Receive(objectOne)
	if pool.NumberOfActiveObjects() != 0 {
		t.Fatal("There should not be active object")
	}
	if pool.NumberOfIdleObjects() != 1 {
		t.Fatal("Actually there should be zero idle objects in pool")
	}

	secondCall := pool.Loan()
	if secondCall.id != createdId {
		t.Fatal(strings.Join([]string{
			"Pool should return same object: we now have ",
			strconv.Itoa(secondCall.id),
			" instead of ",
			strconv.Itoa(createdId),
		}, ""))
	}
	if pool.NumberOfIdleObjects() != 1 {
		t.Fatal("Actually there should be zero idle objects in pool")
	}
}

func TestPool(t *testing.T) {
	pool := InitPool()

	_ = pool.Loan()
	foo := pool.Loan()
	_ = pool.Loan()

	if pool.NumberOfActiveObjects() != 3 {
		t.Fatal("Actually there should be zero idle objects in pool")
	}

	pool.Receive(foo)
	if pool.NumberOfActiveObjects() != 2 {
		t.Fatal("Actually there should be zero idle objects in pool")
	}
}
