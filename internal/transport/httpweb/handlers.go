package webHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	userModel "github.com/sherpaurgen/Tilicho/internal/auth/models"
	"github.com/sherpaurgen/Tilicho/internal/auth/services"
)

type UserService interface {
	GetUserByUsername(context.Context, string) (userModel.User, error)
	CreateUser(context.Context, userModel.User) (string, error)
}

type Handler struct {
	Router  *chi.Mux
	Service *services.UserService
	Server  *http.Server
}

func NewHandler(service *services.UserService) *Handler {
	h := &Handler{
		Router: chi.NewRouter(),

		Service: service,
		Server: &http.Server{
			Addr: ":8080", // Specify the address to listen on
		},
	}

	h.mapRoutes()
	h.Server.Handler = h.Router
	return h
}

func (h *Handler) mapRoutes() {
	// Set up middleware
	h.Router.Use(middleware.Logger)
	h.Router.Use(middleware.Recoverer)

	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	})
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		fmt.Println("Error starting webserver:", err)
		return err
	}
	return nil
}
