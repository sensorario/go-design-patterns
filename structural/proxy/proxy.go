package proxy

import (
	"fmt"
)

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserList []User

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

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

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (u *UserList) addUser(newUser User) {
	*u = append(*u, newUser)
}

func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found\n", id)
}
