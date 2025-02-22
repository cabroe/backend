package main

import (
	"log"
	"net/http"

	"github.com/ptmmeiningen/schichtplaner/app"
)

// @title           Schichtplaner API
// @version         1.0
// @description     REST-API für die Verwaltung von Schichtplänen
// @contact.name    Carsten Bröckert
// @license.name    MIT
// @host            localhost:8080
// @BasePath        /
func main() {
	app, err := app.SetupAndRunApp()
	if err != nil {
		log.Fatalf("Fehler beim Starten der Anwendung: %v", err)
	}

	log.Fatal(app.Listen(":8080"))
}

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "upgrade-insecure-requests")
		next.ServeHTTP(w, r)
	})
}
