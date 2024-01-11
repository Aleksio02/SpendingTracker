package response

import "spending-manager/cmd/spending-manager/model/request"

type SpendingResponse struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Item    request.SpendingRequest `json:"item"`
}
