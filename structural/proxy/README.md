# Proxy

The design pattern usually wraps an object to hide some of its characteristics. The objectives of this patterns are ti hide an object behind the proxy so the features can be hidden, restricted and so on. Again, it provide a new abstraction layer that is easy to work with, and can be changed easily.

In this example we have a database of users.

```go
type User struct {
	ID int32
}

type UserList []User

type UserFinder interface {
	FindUser(id int32) (User, error)
}
```

We dont use a slice of users because declaring a sequence of structs in this way it possibile to implement the interface UserFinder. A slice cant implement an interface.

And here our proxy. A proxy here contains the database and the cache as UserList. The capacity limit the size of cache. All access to the data will be asked to this proxy. The proxy will save in cache new records. When a record will be asked first will be checked the presence in cache, then in database.

```go
type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
}
```

As we said before, we used UserList as sequence of User to implement UserFinder interface. This because our database and our stack are both sequence of user.

```go
func (u *UserListProxy) FindUser(ID int32) (User, error) {
	user, err := u.StackCache.FindUser(ID)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(ID)
	if err == nil {
		fmt.Println("Returning user from database")
		u.addUserToStack(user)
		u.DidLastSearchUsedCache = false
		return user, nil
	}

	return User{}, fmt.Errorf("User not found")
}

func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %s could not be found\n", id)
}
```
