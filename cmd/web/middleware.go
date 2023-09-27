package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a middleware that prevents CSRF attacks.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad is a middleware that loads and saves the session from and to the cookie.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
