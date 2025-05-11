package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/usernamedinesh/learn-go/internal/app"
	"github.com/usernamedinesh/learn-go/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	// http.HandleFunc("/health", healthCheck) // we dont need that because we are using chi here
	r := routes.SetupRoute(app) // pass the app to the chi route

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Println("Starting server on port", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal("Error starting server:", err)
	} else {
		app.Logger.Println("Server started successfully")
	}
}

// we moved this code to the internal/app/app.go
// func healthCheck(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Status is available!\n")
// }
