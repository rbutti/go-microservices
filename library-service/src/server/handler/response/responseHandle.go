package response

import (
	"library-service/util/logger"
)

type RespHandle struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *RespHandle {
	return &RespHandle{logger: logger}
}
func (app *RespHandle) Logger() *logger.Logger {
	return app.logger
}
