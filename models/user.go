package models

import (
	"fmt"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.Id != 0 {
		return User{}, fmt.Errorf("new User must not include id or it must be set to zero")
	}
	user.Id = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserById(id int) (User, error) {
	for _, user := range users {
		if user.Id == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, candidate := range users {
		if candidate.Id == user.Id {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", user.Id)
}

func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID '%v' not found", id)
}
