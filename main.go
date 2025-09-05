package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/somebroyouknow/goHttpFirst/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Application running bro")

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

// r is a  pointer is because it will contain data for our client
// we will see that with use of a middleware, we want to persist the data
// w is something we get from handleFunc that we can modify as we want
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")

}
