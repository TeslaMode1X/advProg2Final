package interfaces

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type RedisCache interface {
	Get(ctx context.Context, userID uuid.UUID) (model.User, error)
	Set(ctx context.Context, user model.User) error
	SetMany(ctx context.Context, users []model.User) error
	Delete(ctx context.Context, userID uuid.UUID) error
}
