package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *Handler) registerMovementRoutes() {
	s.router.Route("/movements", func(r chi.Router) {
		r.Use(s.isAuthenticated)
		r.Get("/", s.getMovementsByUser)
		r.Get("/{movementId}", s.getMovementStats)
		r.Delete("/{movementId}", s.removeMovement)
	})
}

func (s *Handler) getMovementStats(w http.ResponseWriter, r *http.Request) {
	movementIdStr := chi.URLParam(r, "movementId")

	movementId, err := strconv.ParseUint(movementIdStr, 10, 64)
	if err != nil {
		JsonErrorResponse(w, "invalid workoutId parameter", ErrBadRequest, http.StatusBadRequest)
		return
	}

	movement, err := s.svc.GetMovement(movementId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get movement", err, http.StatusInternalServerError)
		return
	}

	if movement.User != s.getCurrentUser(r) {
		JsonErrorResponse(w, "not authorized to remove movement", err, http.StatusUnauthorized)
		return
	}

	movementStats, err := s.svc.GetMovementStats(movementId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get movement stats", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, movementStats, http.StatusOK)
}

func (s *Handler) getMovementsFromWorkout(w http.ResponseWriter, r *http.Request) {
	workoutIdStr := chi.URLParam(r, "workoutId")

	if workoutIdStr == "" {
		JsonErrorResponse(w, "invalid workoutId parameter", ErrBadRequest, http.StatusBadRequest)
		return
	}

	workoutId, err := strconv.ParseUint(workoutIdStr, 10, 64)
	if err != nil {
		JsonErrorResponse(w, "invalid workoutId parameter", ErrBadRequest, http.StatusBadRequest)
		return
	}

	movements, err := s.svc.GetMovementsFromWorkout(workoutId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get movements", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, movements, http.StatusOK)
}

func (s *Handler) getMovementsByUser(w http.ResponseWriter, r *http.Request) {
	userId := s.getCurrentUser(r)

	movements, err := s.svc.GetMovementsByUser(userId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get movements by user", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, movements, http.StatusOK)
}

func (s *Handler) removeMovement(w http.ResponseWriter, r *http.Request) {
	movementIdStr := chi.URLParam(r, "movementId")

	movementId, err := strconv.ParseUint(movementIdStr, 10, 64)
	if err != nil {
		JsonErrorResponse(w, "invalid workoutId parameter", ErrBadRequest, http.StatusBadRequest)
		return
	}

	movement, err := s.svc.GetMovement(movementId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get movement", err, http.StatusInternalServerError)
		return
	}

	if movement.User != s.getCurrentUser(r) {
		JsonErrorResponse(w, "not authorized to remove movement", err, http.StatusUnauthorized)
		return
	}

	err = s.svc.DeleteMovement(movementId)
	if err != nil {
		JsonErrorResponse(w, "couldn't delete movement", err, http.StatusInternalServerError)
		return
	}
}
