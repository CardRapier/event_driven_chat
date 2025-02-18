package chat

import (
	"time"

	"github.com/google/uuid"
)


type ChatMessage struct {
	UserId uuid.UUID
	Message string
	Timestamp time.Time
}