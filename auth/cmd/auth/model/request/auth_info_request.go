package response

type AuthInfoRequest struct {
	Status   int `json:"status"`
	Username int `json:"username"`
	Password int `json:"password"`
	Response any `json:"response"`
}
