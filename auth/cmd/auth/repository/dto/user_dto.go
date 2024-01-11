package dto

type UserDTO struct {
	Id       int    `db="id"`
	Username string `db="username"`
	Password string `db="password"`
	Role     string `db="role"`
}
