package user

import (
	"fmt"
)

type User struct{
	Username string
	Password string
}

func Register(username, password string) User{
	fmt.Println("User registered successfully!")
	return User{Username: username, Password: password}
}

func Login(user User, username, password string) bool{
	return user.Username == username && user.Password==password
}