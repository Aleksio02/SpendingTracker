package response

type SpendingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Item    any    `json:"item"`
}
