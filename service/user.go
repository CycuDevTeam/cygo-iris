package service

import (
	"cygo_iris/repository"

	"strings"
)

func CheckUserExistByUsername(username string) (bool, error) {
	return repository.CheckExistByUsername(username)
}

func CheckUserExistByEmail(email string) (bool, error) {
	return repository.CheckExistByEmail(email)
}

func InsertUser(user *repository.User) error {
	return repository.InsertUser(user)
}

func GetUserByAccount(account string) (user repository.User) {
	if strings.Contains(account, "@") {
		user = repository.GetUserByEmail(account)
	} else {
		user = repository.GetUserByUsername(account)
	}
	return
}

