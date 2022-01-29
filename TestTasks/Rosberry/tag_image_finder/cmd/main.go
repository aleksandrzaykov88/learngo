package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tagimagefinder"
	"tagimagefinder/pkg/repository"
	"tagimagefinder/pkg/service"

	"tagimagefinder/pkg/handler"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil { // get params from config file
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil { // get env variables
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{ // set database params
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)    // work with db
	services := service.NewService(repos)    // business logic
	handlers := handler.NewHandler(services) // work with http

	srv := new(tagimagefinder.Server) // Server initialization
	go func() {                       // Run server in goroutine for gently switch off
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Tag Image Finder API started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Tag Image Finder API shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	db.Close()
}

// initConfig reads the configuration file
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
