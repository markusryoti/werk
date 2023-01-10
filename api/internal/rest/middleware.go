package rest

import (
	"context"
	"errors"
	"net/http"
)

var (
	ErrBadAuthHeader = errors.New("bad auth header")
	ErrInvalidToken  = errors.New("invalid auth token")
)

/*
func (s *Router) isAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		parts := strings.Split(header, " ")

		if len(parts) != 2 {
			JsonErrorResponse(w, "invalid auth header", ErrBadAuthHeader, http.StatusBadRequest)
			return
		}

		token := parts[1]

		uid, err := s.authClient.VerifyToken(token)
		if err != nil {
			JsonErrorResponse(w, "invalid auth token", ErrInvalidToken, http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), "uid", uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
*/

func (s *Router) isAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := s.authClient.GetTokenFromSessionCookie(r)
		if err != nil {
			JsonErrorResponse(w, "not authenticated", err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}