package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/somebroyouknow/goHttpFirst/internal/app"
	"github.com/somebroyouknow/goHttpFirst/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse() //  actually calls the function

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close() // tells application at the very end of execution to go ahead and call this function

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Application running bro on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
