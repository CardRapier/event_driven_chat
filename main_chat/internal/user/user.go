package user

import "github.com/google/uuid"

type CreateUser struct {
	Username string `json:"username"`
}

type ChatUser struct {
	UserId   uuid.UUID `json:"id"`
	Username string `json:"username"`
}