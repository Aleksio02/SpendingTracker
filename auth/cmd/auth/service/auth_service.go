package service

import (
	"auth/cmd/auth/model"
	request "auth/cmd/auth/model/request"
	"auth/cmd/auth/repository"
	"auth/cmd/auth/repository/dto"
	"errors"
)

func GetUserByUsername(user model.User) (model.User, error) {
	foundUser := repository.GetUserByUsername(user.Username)
	if foundUser.Id == 0 {
		return model.User{}, errors.New("user with this username doesn't exist")
	}
	return foundUser, nil
}

func Authorize(request request.AuthRequest) (model.User, error) {
	userDTO := dto.UserDTO{Username: request.Username, Password: request.Password}

	foundUserDTO, err := repository.GetUserByUsernameAndPassword(userDTO)

	if err != nil {
		return model.User{}, errors.New("incorrect username or password")
	}
	return model.User{Username: foundUserDTO.Username, Role: foundUserDTO.Role}, nil
}

func RegisterUser(data request.AuthRequest) (model.User, error) {

	nameUser := model.User{
		Username: data.Username}

	_, err := GetUserByUsername(nameUser)

	if err != nil {
		userInfo := dto.UserDTO{
			Username: data.Username,
			Password: data.Password,
			Role:     data.Role}

		savedUser, sqlerr := repository.SaveUser(userInfo)
		if sqlerr != nil {
			return model.User{}, sqlerr
		}
		return model.User{Username: savedUser.Username, Role: savedUser.Role}, nil
	}

	return model.User{}, errors.New("user with this username is exist")
}
