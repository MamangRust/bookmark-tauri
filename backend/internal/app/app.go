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
	"bookmark-backend/pkg/mapper"
	"bookmark-backend/pkg/upload"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

func Run() {
	log, err := logger.NewLogger()

	if err != nil {
		log.Fatal("err")
	}

	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(runtime.NumCPU() / 2)
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

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Repository: repository,
		Logger:     log,
		Hash:       hashing,
		Token:      token,
		Mapper:     mapper,
	})

	image := upload.NewImage(service.File, service.Folder)

	myhandler := handler.NewHandler(service, token, image)

	// staticDir := http.Dir(viper.GetString("STATIC_DIR"))

	// fs := http.FileServer(http.Dir(staticDir))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("PORT")),
		WriteTimeout: time.Duration(viper.GetInt("WRITE_TIME_OUT")) * time.Second * 10,
		ReadTimeout:  time.Duration(viper.GetInt("READ_TIME_OUT")) * time.Second * 10,

		IdleTimeout: time.Second * 60,
		Handler:     myhandler.Init(),
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Info("Connected to port: " + viper.GetString("PORT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)
}
