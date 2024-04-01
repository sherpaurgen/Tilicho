package webHandler

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
		Router:  chi.NewRouter(),
		Service: service,
	}
	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    ":8080",  // Specify the address to listen on
		Handler: h.Router, // specify routes to use for this web server
	}

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
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c //wait until some value is sent on channel "c" , the value is discarded
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("shutting down webserver")
	return nil
}
