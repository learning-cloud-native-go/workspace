package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"workspace.dev/services/apis/bookapi/app/router"
	"workspace.dev/shared/go/configs"
	"workspace.dev/shared/go/logger"
	"workspace.dev/shared/go/validator"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

//	@title			Book API
//	@version		1.0
//	@description	This is a RESTful Book API with CRUD functionality

//	@contact.name	Dumindu Madunuwan
//	@contact.url	https://learning-cloud-native-go.github.io

//	@license.name	Mozilla Public License Version 2.0
//	@license.url	https://github.com/learning-cloud-native-go/workspace/blob/master/LICENSE

// @servers.url	localhost:8080/v1
func main() {
	c := configs.NewBookAPI()
	l := logger.New(c.Server.Debug)
	v := validator.New()

	var logLevel gormlogger.LogLevel
	if c.DB.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, c.DB.Host, c.DB.Username, c.DB.Password, c.DB.DBName, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		l.Fatal().Err(err).Msg("DB connection start failure")
		return
	}

	r := router.New(l, v, db)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	l.Info().Msgf("Starting server %v", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatal().Err(err).Msg("Server startup failure")
	}
}
