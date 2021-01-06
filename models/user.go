package models

import (
	"errors"
	"fmt"
)

// User is a thing
type User struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
	Age       int
}

var (
	users = []*User{
		&User{ID: 0, FirstName: "Jacob", LastName: "Crawford", Age: 28, Address: "Reginehøjvej 5"},
		&User{ID: 1, FirstName: "Jacob", LastName: "Clausen", Age: 44, Address: "Sønderholt 30"},
		&User{ID: 2, FirstName: "Jacob", LastName: "Nielsen", Age: 25, Address: "Kalmargade 46"},
	}
	nextID = 1
)

// GetUsers is a function that returns the users
func GetUsers() []*User {
	return users
}

// AddUser is a function for adding new users
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not have and id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID is a function for getting users
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if (*u).ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("No user found with id '%v'", id)
}

// GetUsersOfAge is a func
func GetUsersOfAge(age int) []*User {
	usersWithAge := make([]*User, 0)
	for _, u := range users {
		if u.Age == age {
			usersWithAge = append(usersWithAge, u)
		}
	}
	return usersWithAge
}

// UpdateUser is
func UpdateUser(u User) (User, error) {
	for i, cand := range users {
		if cand.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("No user found with id '%v'", u.ID)
}

// RemoveUserByID is
func RemoveUserByID(id int) error {
	for i, cand := range users {
		if cand.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No user found with id '%v'", id)
}
