package jsonapi

import (
	"context"
	"net/http"

	"github.com/ta8i2chi8/go-api-sample/internal/domain/model"
	"github.com/ta8i2chi8/go-api-sample/internal/domain/repository"
)

type PostRepository struct {
	client *jsonApiClient
}

func NewPostRepository(client *jsonApiClient) repository.PostRepository {
	return &PostRepository{client: client}
}

func (r *PostRepository) GetPosts(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post

	// /postsエンドポイントにGETリクエストを送信
	err := r.client.doRequest(ctx, http.MethodGet, "/posts", nil, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
