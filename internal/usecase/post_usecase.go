package usecase

import (
	"context"

	"github.com/ta8i2chi8/go-api-sample/internal/domain/entity"
	"github.com/ta8i2chi8/go-api-sample/internal/domain/repository"
)

type PostUsecase interface {
	GetPosts(ctx context.Context) ([]entity.Post, error)
}

type postUsecase struct {
	postRepository repository.PostRepository
}

func NewPostUsecase(postRepository repository.PostRepository) *postUsecase {
	return &postUsecase{
		postRepository: postRepository,
	}
}
func (p *postUsecase) GetPosts(ctx context.Context) ([]entity.Post, error) {
	return p.postRepository.GetPosts(ctx)
}
