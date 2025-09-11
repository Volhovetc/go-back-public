package main

import (
	"fmt"
	"go-back/configs"
	"go-back/internal/auth"
	"go-back/internal/link"
	"go-back/internal/stat"
	"go-back/internal/user"
	"go-back/pkg/dbPostgresComposable"
	"go-back/pkg/event"
	"go-back/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := dbPostgresComposable.NewDb(conf)
	mux := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repository
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	// Handler
	auth.NewAuthHandler(mux, auth.AuthHandlerDeps{
		AuthConfig:  &conf.Auth,
		AuthService: authService,
	})
	link.NewLinkHandler(mux, link.LinkHandlerDeps{
		Config:         &conf.Auth,
		EventBus:       eventBus,
		LinkRepository: linkRepository,
	})

	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(mux),
	}
	go statService.AddClick()
	fmt.Println("Starting server on port 8081...")
	server.ListenAndServe()
}
