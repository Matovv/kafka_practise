package http

import (
	"local/service1/internal/pkg/configs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	rootRouter chi.Router
}

func NewRouter(config *configs.Config) *Router {
	return &Router{}
}

func (r *Router) Handler() http.Handler {
	return r.rootRouter
}