package app

import (
	"bookmark-backend/internal/handler"
	"bookmark-backend/internal/repository"
	"bookmark-backend/internal/service"
	"bookmark-backend/pkg/auth"
	"bookmark-backend/pkg/database/postgres"
	"bookmark-backend/pkg/dotenv"
	"bookmark-backend/pkg/hash"
	"bookmark-backend/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"
)

func Run() {
	log, err := logger.NewLogger()

	if err != nil {
		log.Fatal("err")
	}

	err = dotenv.Viper()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewClient(*log)

	if err != nil {
		log.Fatal(err.Error())
	}

	hashing := hash.NewHashingPassword()

	repository := repository.NewRepositories(db)

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		log.Fatal(err.Error())
	}

	service := service.NewService(service.Deps{
		Repository: repository,
		Logger:     *log,
		Hash:       *hashing,
		Token:      token,
	})

	myhandler := handler.NewHandler(service, token)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("PORT")),
		IdleTimeout:  120 * 1000,
		WriteTimeout: 15 * 60 * 1000,
		ReadTimeout:  15 * 60 * 1000,
		Handler:      myhandler.Init(),
	}

	go func() {
		if err := serve.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Info("Server running on port :8080")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)
}
