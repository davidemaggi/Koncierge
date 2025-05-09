package container

import (
	"github.com/davidemaggi/koncierge/internal/logger"
)

type Services struct {
	Logger *logger.Logger
}

var App *Services

func Init(isVerbose bool) {
	App = &Services{
		Logger: logger.NewLogger(isVerbose),
	}
}
