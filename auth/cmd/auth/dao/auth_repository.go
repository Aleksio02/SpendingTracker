package dao

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/dao/dto"
	"auth/cmd/auth/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func GetUserByUsername(username string) model.User {
	connection, _ := config.CreateDatabaseConnection()
	defer connection.Close(context.Background())

	row, _ := connection.Query(context.Background(), "SELECT u.id as id, username, password, r.name as role FROM users u INNER JOIN ref_user_role r ON r.id = u.role_id WHERE upper(username) = upper('"+username+"')")
	foundUser, err := pgx.CollectOneRow(row, pgx.RowToStructByName[dto.UserDTO])
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return model.User{Id: foundUser.Id, Username: foundUser.Username, Role: foundUser.Role}
}

func SaveUser(data dto.UserDTO) (dto.UserDTO, error) {
	//Сохраняем экземпляр DTO в БД(PostgreSQL) и возвращаем модель пользователя

	connection, _ := config.CreateDatabaseConnection()
	defer connection.Close(context.Background())

	_, err := connection.Query(context.Background(), "INSERT INTO users(username, password, role_id) VALUES ('"+data.Username+"', '"+data.Password+"', (select id from ref_user_role where name = '"+data.Role+"'))")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return dto.UserDTO{}, errors.New("error occurred while executing SQL query")
	}

	return data, nil
}
