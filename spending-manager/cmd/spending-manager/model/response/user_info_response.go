package response

type UserInfoResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
