package dto

type UserDto struct {
	Username string `json:"username" binding:"required"`
}
