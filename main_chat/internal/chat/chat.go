package chat

import (
	"github.com/google/uuid"
)


type CreateChatMessage struct {
	UserId uuid.UUID `json:"userId"`
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

type ChatMessage struct {
	UserId uuid.UUID `json:"userId"`
	MessageId uuid.UUID `json:"messageId"`
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}