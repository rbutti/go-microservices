package main

import (
	"fmt"
	"library-service/config"
	handler "library-service/server/handler/response"
	"library-service/server/router"
	lr "library-service/util/logger"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	responseHandler := handler.New(logger)

	appRouter := router.New(responseHandler)

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
