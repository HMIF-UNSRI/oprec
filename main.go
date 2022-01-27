package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"oprec/controller"
	"oprec/middleware"
	"oprec/pkg/helper"
	"oprec/pkg/mariadb"
	"oprec/repository"
	"oprec/usecase"
	"os"
)

func main() {
	conn := mariadb.NewConnection()
	validate := validator.New()

	participantRepository := repository.NewParticipantRepositoryImpl(conn)
	participantUsecase := usecase.NewParticipantUsecaseImpl(participantRepository, validate)

	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	helper.PanicIfErr(err)
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(multiWriter)
	loggerMiddleware := middleware.LoggerMiddleware(logger)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	controller.RegisterParticipantController(mux, participantUsecase)

	loggedMux := loggerMiddleware(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}
	logger.Fatalln(server.ListenAndServe())
}
