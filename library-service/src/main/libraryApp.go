package main

import (
	"fmt"
	"library-service/config"
	dbConn "library-service/database/orm"
	handler "library-service/server/handler/response"
	"library-service/server/router"
	lr "library-service/util/logger"
	"net/http"
	"time"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	time.Sleep(15 * time.Second)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	if appConf.Debug {
		db.LogMode(true)
	}

	application := handler.New(logger, db)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
