package repository

import (
	"context"
	"errors"
	"fmt"
	"spending-manager/cmd/spending-manager/config"
	"spending-manager/cmd/spending-manager/repository/dto"

	"github.com/jackc/pgx/v5"
)

func AddSpentItem(spentItem dto.SpentItemDTO) (dto.SpentItemDTO, error) {
	connection, _ := config.CreateDatabaseConnection()
	defer connection.Close(context.Background())

	fmt.Println(spentItem.Description)
	fmt.Println(spentItem.CreateDate)
	fmt.Println(spentItem.Spent)
	fmt.Println(spentItem.Category)
	fmt.Println(spentItem.Username)

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

	queryRow := "select id, description, create_date, spent, (select name from ref_spending_category sc where sc.id = s.category_id) as category, '" + username + "' as username from spending s where user_id = (select id from users where username = '" + username + "')"

	rows, err := connection.Query(context.Background(), queryRow)
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.SpentItemDTO])
	if err != nil {
		return []dto.SpentItemDTO{}, errors.New("error occurred while executing SQL query")
	}

	return items, nil
}
