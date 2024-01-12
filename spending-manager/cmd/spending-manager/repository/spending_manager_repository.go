package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"spending-manager/cmd/spending-manager/config"
	"spending-manager/cmd/spending-manager/repository/dto"
)

func AddSpentItem(spentItem dto.SpentItemDTO) (dto.SpentItemDTO, error) {
	connection, _ := config.CreateDatabaseConnection()
	defer connection.Close(context.Background())

	queryRow := fmt.Sprintf("insert into spending(description, create_date, spent, category_id, user_id) values ('%s', TO_DATE('%s', 'YYYY-MM-DDTHH24:MI:SSZ'), %f, (SELECT id FROM ref_spending_category WHERE name = '%s'),(SELECT id FROM users WHERE username = '%s')) RETURNING id",
		spentItem.Description, spentItem.CreateDate, spentItem.Spent, spentItem.Category, spentItem.Username)

	row, err := connection.Query(context.Background(), queryRow)
	id, err := pgx.CollectOneRow(row, pgx.RowTo[int32])

	if err != nil {
		return dto.SpentItemDTO{}, errors.New("error occurred while executing SQL query")
	}

	spentItem.Id = id
	return spentItem, nil
}

func GetSpendingsByUsername(username string) ([]dto.SpentItemDTO, error) {
	connection, _ := config.CreateDatabaseConnection()
	defer connection.Close(context.Background())

	queryRow := "select * from spending where user_id = (select id from users where username = '" + username + "')"

	row, err := connection.Query(context.Background(), queryRow)
	items, err := pgx.CollectRows(row, pgx.RowToStructByName[dto.SpentItemDTO])

	if err != nil {
		return []dto.SpentItemDTO{}, errors.New("error occurred while executing SQL query")
	}

	return items, nil
}
