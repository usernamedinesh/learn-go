package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/usernamedinesh/learn-go/internal/app"
)

// here we are usign chi like express
func SetupRoute(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	return r
}
