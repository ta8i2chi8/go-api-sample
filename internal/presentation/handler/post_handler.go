package handler

import (
	"log/slog"
	"net/http"

	"github.com/ta8i2chi8/go-api-sample/internal/presentation/common"
	"github.com/ta8i2chi8/go-api-sample/internal/usecase"
)

type PostHandler struct {
	PostUsecase usecase.PostUsecase
}

func NewPostHandler(postUsecase usecase.PostUsecase) *PostHandler {
	return &PostHandler{
		PostUsecase: postUsecase,
	}
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.PostUsecase.GetPosts(ctx)
	if err != nil {
		common.WriteErrorResponse(ctx, w, http.StatusInternalServerError, err.Error())
		slog.Error("failed to get posts", "err", err)
		return
	}

	common.WriteSuccessResponse(ctx, w, posts)
}
