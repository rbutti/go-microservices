package handler

import (
	"github.com/jinzhu/gorm"
	"library-service/util/logger"
)

type Handler struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(
	logger *logger.Logger,
	db *gorm.DB,
) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

func (app *Handler) Logger() *logger.Logger {
	return app.logger
}
