package models

import (
	"errors"
	"fmt"
)

// User defined fields
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers function to return them
func GetUsers() []*User {
	return users
}

// AddUser function that checks if id was passed and if so it's 0
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns the user when their ID is passed
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser will update user by ID, error if ID isn't found
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID will remove the user of ID given, error if not found
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
