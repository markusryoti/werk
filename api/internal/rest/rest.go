package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/markusryoti/werk/internal/auth"
	"github.com/markusryoti/werk/internal/service"
	"go.uber.org/zap"
)

type Router struct {
	router     *chi.Mux
	logger     *zap.SugaredLogger
	authClient *auth.Client
	repo       service.Repository
}

func NewRouter(router *chi.Mux, repo service.Repository, logger *zap.SugaredLogger, authClient *auth.Client) *Router {
	s := &Router{
		router:     router,
		logger:     logger,
		authClient: authClient,
		repo:       repo,
	}

	s.registerRoutes()

	return s
}

func (s *Router) Get() http.Handler {
	return s.router
}

func (s *Router) registerRoutes() {
	s.setCors()

	s.router.Get("/", s.indexHandler)

	s.router.Route("/workouts", func(r chi.Router) {
		r.Use(s.isAuthenticated)
		r.Get("/", s.getAllWorkouts)
		r.Post("/add", s.addWorkoutHandler)
		r.Post("/{workoutId}/addMovement", s.addMovementHandler)
		r.Get("/{workoutId}/workoutMovements", s.getMovementsFromWorkout)
	})
}

func (s *Router) setCors() {
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func (s *Router) indexHandler(w http.ResponseWriter, r *http.Request) {
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
