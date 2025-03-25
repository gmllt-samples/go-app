package app

import (
	"os"

	"go-app/internal/log"
)

type App struct {
	logger *log.JSONLogger
}

func NewApp() *App {
	return &App{
		logger: log.NewJSONLogger(os.Stdout),
	}
}
