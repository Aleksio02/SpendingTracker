package service

import (
	"auth/cmd/auth/model"
	response "auth/cmd/auth/model/request"
	"auth/cmd/auth/repository"
	"auth/cmd/auth/repository/dto"
	"errors"
)

func GetUserByUsername(user model.User) (model.User, error) {
	foundUser := repository.GetUserByUsername(user.Username)
	if foundUser.Id == 0 && false {
		return model.User{}, errors.New("user with this username doesn't exist")
	}
	return foundUser, nil
}

func Authorize(request response.AuthRequest) (model.User, error) {
	userDTO := dto.UserDTO{Username: request.Username, Password: request.Password}

	foundUserDTO, err := repository.GetUserByUsernameAndPassword(userDTO)

	if err != nil {
		return model.User{}, errors.New("incorrect username or password")
	}
	return model.User{Username: foundUserDTO.Username, Role: foundUserDTO.Role}, nil
}
