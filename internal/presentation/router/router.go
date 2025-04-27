package router

import (
	"net/http"

	"github.com/ta8i2chi8/go-api-sample/internal/infrastructure/jsonapi"
	"github.com/ta8i2chi8/go-api-sample/internal/presentation/handler"
	"github.com/ta8i2chi8/go-api-sample/internal/usecase"
)

func New() (http.Handler, error) {
	mux := http.NewServeMux()

	jsonApiClient := jsonapi.NewJsonApiClient("https://jsonplaceholder.typicode.com")
	postRepository := jsonapi.NewPostRepository(jsonApiClient)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postHandler := handler.NewPostHandler(postUsecase)

	mux.HandleFunc("/health", handler.HealthCheckHandler)
	mux.HandleFunc("/posts", postHandler.GetPosts)

	return mux, nil
}
