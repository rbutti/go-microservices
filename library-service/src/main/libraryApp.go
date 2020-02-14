package main

import (
	"fmt"
	"library-service/config"
	"library-service/server/handler"
	"library-service/server/router"
	"library-service/util/logger"
	"net/http"
	"time"
)

func main() {

	//Wait for DB to get loaded
	time.Sleep(30 * time.Second)

	appConf := config.GetAppConfig()
	appLogger := logger.New(appConf.Debug)

	//get DB connection
	db, err := config.GetORMConfig(appConf)
	if err != nil {
		appLogger.Fatal().Err(err).Msg("")
		return
	}
	if appConf.Debug {
		db.LogMode(true)
	}

	//initiate router and handler
	appHandler := handler.New(appLogger, db)
	appRouter := router.New(appHandler)

	//initiate the server
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	appLogger.Info().Msgf("Server starting... %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		appLogger.Fatal().Err(err).Msg("Server startup failed")
	}
	//notify success
	appLogger.Info().Msgf("Application started Successfully!! %v", address)
}
