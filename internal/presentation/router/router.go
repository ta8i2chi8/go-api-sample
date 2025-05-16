package router

import (
	"net/http"

	"github.com/ta8i2chi8/go-api-sample/internal/infra/jsonapi"
	"github.com/ta8i2chi8/go-api-sample/internal/presentation/handler"
	"github.com/ta8i2chi8/go-api-sample/internal/presentation/middleware"
	"github.com/ta8i2chi8/go-api-sample/internal/usecase"
)

type customHandler struct {
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func newCustomHandler() *customHandler {
	return &customHandler{}
}

func (ch *customHandler) use(mw func(http.HandlerFunc) http.HandlerFunc) {
	ch.middlewares = append(ch.middlewares, mw)
}

func (ch *customHandler) wrap(handler http.HandlerFunc) http.HandlerFunc {
	for _, v := range ch.middlewares {
		handler = v(handler)
	}

	return handler
}

func New() (http.Handler, error) {
	mux := http.NewServeMux()
	ch := newCustomHandler()

	mux.HandleFunc("/health", handler.HealthCheckHandler)

	ch.use(middleware.CheckBearerToken)

	jsonApiClient := jsonapi.NewJsonApiClient("https://jsonplaceholder.typicode.com")
	postRepository := jsonapi.NewPostRepository(jsonApiClient)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postHandler := handler.NewPostHandler(postUsecase)
	mux.HandleFunc("/posts", ch.wrap(postHandler.GetPosts))

	return mux, nil
}
