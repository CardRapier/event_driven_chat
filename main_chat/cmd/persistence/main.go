package persistence

import (
	"chat/internal/chat"
	"chat/internal/config"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/segmentio/kafka-go"
)

func StartPersistenceService() {
	redpanda := startRedpandaReader()
	redpandaChatConsumer := chat.NewRedpandaChatConsumer(redpanda)

	go redpandaChatConsumer.ConsumeChatMessages()
}

func startScylla() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1") // Replace with your Scylla/Cassandra host(s)
	cluster.Keyspace = "mandcondor_chat"     // Replace with your keyspace
	cluster.Consistency = gocql.Quorum
	// Optionally, adjust timeouts and other settings:
	cluster.Timeout = 10 * time.Second

	// Create a session. You may wish to handle errors differently in production.
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB: %v", err)
	}

	return session
}

func startRedpandaReader() *kafka.Reader {
	brokerAddress := config.Envs.RedpandaUrl // Update with your Redpanda broker address
	topic := "chat-messages"
	groupID := "persistence-group"

	// Create a new reader with the topic and the broker address
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		GroupID:  groupID, // Set the GroupID here
		MinBytes: 10e3,    // 10KB
		MaxBytes: 1e6,     // 1MB
	})

	log.Println("Consumer - Connected")

	return reader
}
