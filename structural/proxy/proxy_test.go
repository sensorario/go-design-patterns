package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}
	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCapacity: 2,
		StackCache:    UserList{},
	}

	knownIDs := [3]int32{
		someDatabase[3].ID,
		someDatabase[4].ID,
		someDatabase[5].ID,
	}

	t.Run("FindUser = Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesnt match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned frmo an emtpy cache")
		}
	})

	t.Run("FindUser = One user ask for same user", func(t *testing.T) {
		_, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if !proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned frmo an emtpy cache")
		}
	})
}
