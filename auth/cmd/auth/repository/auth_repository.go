package repository

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/model"
	"auth/cmd/auth/repository/dto"
	"context"
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
		os.Exit(1)
	}
	return model.User{Id: foundUser.Id, Username: foundUser.Username, Role: foundUser.Role}
}
