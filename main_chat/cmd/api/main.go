package api

import (
	"chat/internal/chat"
	"chat/internal/config"
	"chat/internal/user"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Handlers struct {
	ChatHandler *chat.ChatHandler
	UserHandler *user.UserHandler
}

type Services struct {
	ChatService *chat.ChatService
	UserService *user.UserService
}

type Dependencies struct {
	Redpanda             *kafka.Writer
	RedpandaChatProducer chat.ChatProducer
	RedisUserRepository  user.UserRepository
}

func StartApi() {
	app := fiber.New()

	// scyllaSession := startScylla()
	redpanda := startRedPanda()
	redis := startRedis()

	dependencies := createDependencies(redis, redpanda)
	services := createServices(dependencies)
	handlers := createHandlers(services)
	setupRoutes(app, handlers)

	handleServer(app, dependencies)
}

func setupRoutes(app *fiber.App, handlers *Handlers) {
	api := app.Group("/api")
	v1 := api.Group("v1")

	handlers.UserHandler.StartRouting(v1)
	handlers.ChatHandler.StartRouting(v1)
}

func createDependencies(redis *redis.Client, redpanda *kafka.Writer) *Dependencies {
	return &Dependencies{
		Redpanda:             redpanda,
		RedpandaChatProducer: chat.NewRedpandaChatProducer(redpanda),
		RedisUserRepository:  user.NewRedisRepository(redis),
	}
}

func createServices(repositories *Dependencies) *Services {
	return &Services{
		ChatService: chat.NewChatService(repositories.RedpandaChatProducer),
		UserService: user.NewUserService(repositories.RedisUserRepository),
	}
}

func createHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler: user.NewUserHandler(services.UserService),
		ChatHandler: chat.NewChatHandler(services.ChatService),
	}
}

func startRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Envs.RedisUrl,
		Password: config.Envs.RedisPass,
		DB:       0,
	})

	// context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection error: %v", err)
	}

	log.Printf("Redis connected successfully: %s", pong)

	return redisClient
}

func tryRedpandaConn() error {
	conn, err := kafka.Dial("tcp", config.Envs.RedpandaUrl)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Println("Connected successfully to Redpanda broker!")
	return nil
}

func startRedPanda() *kafka.Writer {
	err := tryRedpandaConn()

	if err != nil {
		log.Fatalf("failed to connect to Redpanda broker: %v", err)
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.Envs.RedpandaUrl}, // Redpanda broker(s)
		Topic:   "chat-messages",
		// Balancer: &kafka.Hash{}, // Use Hash to ensure messages with the same key go to the same partition.x
	})

	return writer
}

func closeRedpanda(redpanda *kafka.Writer) {
	if err := redpanda.Close(); err != nil {
		log.Fatalf("failed to close writer: %v", err)
	}
	log.Println("Redpanda closed")
}

func handleServer(app *fiber.App, dependencies *Dependencies) {
	// Create a proper context for shutdown
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	// Run server in goroutine
	go func() {
		if err := app.Listen(":" + config.Envs.Port); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	log.Println("Shutting down server...")

	closeRedpanda(dependencies.Redpanda)

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
