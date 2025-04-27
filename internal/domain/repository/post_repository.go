package repository

import (
	"context"

	"github.com/ta8i2chi8/go-api-sample/internal/domain/entity"
)

type PostRepository interface {
	GetPosts(ctx context.Context) ([]entity.Post, error)
}
