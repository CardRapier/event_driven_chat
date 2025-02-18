package user

import "github.com/google/uuid"


type UserRepository interface {
	GetUserByID(userId *uuid.UUID) (*ChatUser, error)
	SaveUser(user *ChatUser) error
}
