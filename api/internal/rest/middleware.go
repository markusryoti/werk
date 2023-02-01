package rest

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

var (
	ErrBadAuthHeader = errors.New("bad auth header")
	ErrInvalidToken  = errors.New("invalid auth token")
)

type userCtxKey string

func (u userCtxKey) String() string {
	return string(u)
}

var userKey = userCtxKey("uid")

func (s *Handler) isAuthenticated(next http.Handler) http.Handler {
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
			return
		}

		ctx := context.WithValue(r.Context(), userKey, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Handler) getCurrentUser(r *http.Request) string {
	uid, _ := r.Context().Value(userKey).(string)
	return uid
}
