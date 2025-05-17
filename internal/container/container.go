package container

import (
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/logger"
	"os"
)

type Services struct {
	Logger *logger.Logger
}

var App *Services

func Init(isVerbose bool) {

	lg := logger.NewLogger(isVerbose)

	db.Init()
	err := db.Migrate()

	if err != nil {
		lg.Error("Cannot Instantiate Database", err)
		os.Exit(1)
	}

	App = &Services{
		Logger: lg,
	}

}
