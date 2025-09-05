package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// as app is the memory in an application, what happens if you remmove the &  during its declaration
	app := &Application{
		Logger: logger,
	}

	return app, nil // nil are valid error types so we put it here

}
