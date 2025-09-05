package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/somebroyouknow/goHttpFirst/internal/api"
	"github.com/somebroyouknow/goHttpFirst/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDb, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// space for stores

	// here are my handlers

	workoutHandler := api.NewWorkoutHandler()

	// as app is the memory in an application, what happens if you remmove the &  during its declaration
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDb,
	}

	return app, nil // nil are valid error types so we put it here

}

// r is a  pointer is because it will contain data for our client
// we will see that with use of a middleware, we want to persist the data
// w is something we get from handleFunc that we can modify as we want
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthcheck successful ")

}
