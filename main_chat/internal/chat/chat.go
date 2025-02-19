package chat

import (
	"github.com/google/uuid"
)


type ChatMessage struct {
	UserId uuid.UUID `json:"userId"`
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}