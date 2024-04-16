package types

import (
	"github.com/google/uuid"
)

type UserInput struct {
	Login string `json:"login" binding:"required" fake:"{regex:^[a-z]{1,15}$}"`
	Name *string `json:"name" fake:"{regex:^[a-z]{1,15}$}"`
	Phone string `json:"phone" binding:"required" validate:"phone" fake:"{regex:7910\\d{7}$}"`
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
