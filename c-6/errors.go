package main

import (
	"errors"
	"fmt"
)

type User struct {
	name string
}

func main() {
	user, err := getUser(User{name: "demarcus"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}

func getUser(user User) (User, error) {
	if user.name != "" {
		return user, nil
	} else {
		return user, errors.New("Something went wrong!")
	}
}
