package types

import "github.com/google/uuid"

type UserInput struct {
	Login string `json:"login" binding:"required"`
	Name *string `json:"name"`
	Phone string `json:"phone" binding:"required" validate:"phone"`
}

type UserInputUpd struct {
	Name        *string `json:"name,omitempty"`
	Discription *string `json:"discription,omitempty"`
	ID          uuid.UUID
}

type UserUpdReq struct {
    Name        *string `json:"name" binding:"omitempty"`
    Discription *string `json:"discription" binding:"omitempty"`
	ID          *uuid.UUID `json:"id" binding:"omitempty"`
}
