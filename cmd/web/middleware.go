package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	fmt.Println("NoSurf Started")
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteStrictMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves a session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
