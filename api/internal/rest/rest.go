package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/markusryoti/werk/internal/auth"
	"github.com/markusryoti/werk/internal/service"
	"go.uber.org/zap"
)

var (
	ErrBadRequest    = errors.New("bad request")
	ErrNotAuthorized = errors.New("not authorized")
)

type Handler struct {
	router     *chi.Mux
	logger     *zap.SugaredLogger
	authClient *auth.Client
	svc        *service.Service
}

func NewHandler(router *chi.Mux, svc *service.Service, logger *zap.SugaredLogger, authClient *auth.Client) *Handler {
	s := &Handler{
		router:     router,
		logger:     logger,
		authClient: authClient,
		svc:        svc,
	}

	s.registerRoutes()

	return s
}

func (s *Handler) Get() http.Handler {
	return s.router
}

func (s *Handler) registerRoutes() {
	s.setCors()

	s.router.Get("/", s.indexHandler)

	s.registerWorkoutRoutes()
	s.registerMovementRoutes()
	s.registerMovementSetRoutes()
}

func (s *Handler) setCors() {
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func (s *Handler) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hellooo"))
}

func JsonResponse(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&body)
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func JsonErrorResponse(w http.ResponseWriter, message string, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := ErrorResponse{
		Message: message,
		Error:   err.Error(),
	}

	json.NewEncoder(w).Encode(&res)
}

type MessageResponse struct {
	Message string `json:"message"`
}

func JsonMessageResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := MessageResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(&res)
}
