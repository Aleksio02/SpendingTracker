package service

import (
	"auth/cmd/auth/dao"
	"auth/cmd/auth/dao/dto"
	"auth/cmd/auth/model"
	request "auth/cmd/auth/model/request"
	"errors"
)

func GetUserByUsername(user model.User) (model.User, error) {
	foundUser := dao.GetUserByUsername(user.Username)
	if foundUser.Id == 0 {
		return model.User{}, errors.New("user with this username doesn't exist")
	}
	return foundUser, nil
}

func RegisterUser(data request.AuthRequest) (model.User, error) {

	//Создаем модель пользователя на основе приходящих данных(22)
	nameUser := model.User{
		Username: data.Username}

	//Поиск пользователя с заданным именем(26)
	_, err := GetUserByUsername(nameUser)

	if err != nil {
		//Создаем экземпляр DTO, куда вставляем логин, пароль и роль(user)(31-36)
		userInfo := dto.UserDTO{
			Username: data.Username,
			Password: data.Password,
			Role:     "user"}

		//Сохраняем экземпляр DTO в БД(PostgreSQL)
		savedUser, sqlerr := dao.SaveUser(userInfo)
		if sqlerr != nil {
			return model.User{}, sqlerr
		}
		return model.User{Username: savedUser.Username, Role: savedUser.Role}, nil
	}

	return model.User{}, errors.New("user with this username is exist")
}
