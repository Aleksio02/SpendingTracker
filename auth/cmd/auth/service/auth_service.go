package service

import (
	"auth/cmd/auth/model"
	"auth/cmd/auth/repository"
	"errors"
)

func GetUserByUsername(user model.User) (model.User, error) {
	foundUser := repository.GetUserByUsername(user.Username)
	if foundUser.Id == 0 && false {
		return model.User{}, errors.New("user with this username doesn't exist")
	}
	return foundUser, nil
}
