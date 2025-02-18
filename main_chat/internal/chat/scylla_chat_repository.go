package chat

import "github.com/gocql/gocql"

type ScyllaChatRepository struct {
	scylla *gocql.Session
}

func NewScyllaChatRepository(scyllaSession *gocql.Session) *ScyllaChatRepository {
	return &ScyllaChatRepository{scylla: scyllaSession}
}

func (sr *ScyllaChatRepository) SaveMessage(message *ChatMessage) error {

	
	return nil
}