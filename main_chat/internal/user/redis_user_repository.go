package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{client: redisClient}
}

func (r *RedisUserRepository) GetUserByID(userID *uuid.UUID) (*ChatUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get user data from Redis
	val, err := r.client.Get(ctx, userID.String()).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user from Redis: %w", err)
	}

	// Deserialize JSON to ChatUser
	var user ChatUser
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	return &user, nil
}

func (r *RedisUserRepository) SaveUser(user *ChatUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Serialize user data to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	// Store user data with expiration (e.g., 24 hours)
	err = r.client.Set(ctx, user.UserId.String(), userJSON, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to save user to Redis: %w", err)
	}

	return nil
}
