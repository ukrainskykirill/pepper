package types

type UserInput struct {
	Login string `json:"login" binding:"required"`
	Name string `json:"name"`
	Phone string `json:"phone" binding:"required" validate:"phone"`
}