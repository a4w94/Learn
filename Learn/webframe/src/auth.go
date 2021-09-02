package userlogin

import (
	"errors"
	"fmt"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"test": "test",
	}
}

func CheckUserExist(username string) bool {
	_, isExist := UserData[username]

	return isExist
}

func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New(("password is not correct "))
	}
}

func Auth(username string, password string) error {
	if isExist := CheckUserExist(username); isExist {
		fmt.Println(CheckPassword(UserData[username], password))
		return CheckPassword(UserData[username], password)
	} else {
		return errors.New("user is no exist")
	}
}
