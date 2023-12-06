package response

type AuthInfoRequest struct {
	Username int `json:"username"`
	Password int `json:"password"`
	Response any `json:"response"`
}
