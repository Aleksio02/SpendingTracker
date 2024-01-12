package service

import (
	"spending-manager/cmd/spending-manager/model/request"
	"spending-manager/cmd/spending-manager/model/response"
	"spending-manager/cmd/spending-manager/repository"
	"spending-manager/cmd/spending-manager/repository/dto"
)

func AddSpentItem(request request.SpendingRequest, username string) response.SpendingResponse {
	spentItemDTO := dto.SpentItemDTO{Category: request.Category, Description: request.Description, Spent: request.Spent, CreateDate: request.CreateDate, Username: username}

	var err error
	spentItemDTO, err = repository.AddSpentItem(spentItemDTO)

	if err != nil {
		return response.SpendingResponse{
			Status:  500,
			Message: err.Error(),
		}
	}

	request.Id = spentItemDTO.Id
	return response.SpendingResponse{
		Status:  200,
		Message: "Spent item added successfully",
		Item:    request,
	}
}
