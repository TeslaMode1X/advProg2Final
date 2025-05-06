package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/redis"
	"github.com/gofrs/uuid"
	goredis "github.com/redis/go-redis/v9"
)

const (
	keyPrefix = "user:%s"
)

type Client struct {
	client *redis.Client
	ttl    time.Duration
}

func NewClient(client *redis.Client, ttl time.Duration) *Client {
	return &Client{
		client: client,
		ttl:    ttl,
	}
}

func (c *Client) Set(ctx context.Context, user model.User) error {
	// Convert User to UserRedis
	userRedis := model.UserRedis{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	// Marshal UserRedis instead of User
	data, err := json.Marshal(userRedis)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	return c.client.Unwrap().Set(ctx, c.key(user.ID), data, c.ttl).Err()
}

func (c *Client) SetMany(ctx context.Context, users []model.User) error {
	pipe := c.client.Unwrap().Pipeline()
	for _, user := range users {
		data, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("failed to marshal user: %w", err)
		}
		pipe.Set(ctx, c.key(user.ID), data, c.ttl)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to set many users: %w", err)
	}
	return nil
}

func (c *Client) Get(ctx context.Context, userID uuid.UUID) (model.User, error) {
	data, err := c.client.Unwrap().Get(ctx, c.key(userID)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return model.User{}, nil // not found
		}
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	// Unmarshal to UserRedis first
	var userRedis model.UserRedis
	err = json.Unmarshal(data, &userRedis)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	// Convert UserRedis back to User
	user := model.User{
		ID:        userRedis.ID,
		Username:  userRedis.Username,
		Password:  userRedis.Password,
		Email:     userRedis.Email,
		CreatedAt: userRedis.CreatedAt,
		UpdatedAt: userRedis.UpdatedAt,
		DeletedAt: userRedis.DeletedAt,
	}

	return user, nil
}

func (c *Client) Delete(ctx context.Context, userID uuid.UUID) error {
	return c.client.Unwrap().Del(ctx, c.key(userID)).Err()
}

func (c *Client) key(id uuid.UUID) string {
	return fmt.Sprintf(keyPrefix, id.String())
}
