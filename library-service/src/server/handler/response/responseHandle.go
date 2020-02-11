package response

import (
	"github.com/jinzhu/gorm"
	"library-service/util/logger"
)

type RespHandle struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(
	logger *logger.Logger,
	db *gorm.DB,
) *RespHandle {
	return &RespHandle{
		logger: logger,
		db:     db,
	}
}

func (app *RespHandle) Logger() *logger.Logger {
	return app.logger
}
