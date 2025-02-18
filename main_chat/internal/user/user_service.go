package user

import (
	"errors"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (cs *UserService) CreateUser(newUser *CreateUser) (*ChatUser, error) {
	user := &ChatUser{
		UserId:   uuid.New(),
		Username: newUser.Username,
	}
	err := cs.userRepository.SaveUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (cs *UserService) CheckIfUserExists(userId *uuid.UUID) error {
	user, err := cs.userRepository.GetUserByID(userId)

	if err != nil {
		return err
	}

	if user.UserId.String() != userId.String() {
		return errors.New("user not equal")
	}

	return nil
}
