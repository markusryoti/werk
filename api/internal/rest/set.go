package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/markusryoti/werk/internal/types"
)

func (s *Handler) registerMovementSetRoutes() {
	s.router.Route("/sets", func(r chi.Router) {
		r.Use(s.isAuthenticated)
		r.Delete("/{setId}", s.removeSet)
	})
}

type AddSetRequest struct {
	Reps   uint8 `json:"reps"`
	Weight int   `json:"weight"`
}

func (s *Handler) addSet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody AddSetRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		JsonErrorResponse(w, "invalid message body", err, http.StatusBadRequest)
		return
	}

	movementIdStr := chi.URLParam(r, "movementId")
	movementId, err := strconv.ParseUint(movementIdStr, 10, 64)

	userId := s.getCurrentUser(r)

	err = s.svc.AddMovementSet(movementId, types.Set{
		Reps:   reqBody.Reps,
		Weight: reqBody.Weight,
		User:   userId,
	})

	if err != nil {
		JsonErrorResponse(w, "couldn't add a new set", err, http.StatusInternalServerError)
		return
	}
}

func (s *Handler) removeSet(w http.ResponseWriter, r *http.Request) {
	setIdStr := chi.URLParam(r, "setId")

	setId, err := strconv.ParseUint(setIdStr, 10, 64)
	if err != nil {
		JsonErrorResponse(w, "invalid delete request", err, http.StatusBadRequest)
		return
	}

	set, err := s.svc.GetMovementSet(setId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get movement set", err, http.StatusInternalServerError)
		return
	}

	if set.User != s.getCurrentUser(r) {
		JsonErrorResponse(w, "not authorized to remove set", err, http.StatusUnauthorized)
		return
	}

	err = s.svc.DeleteMovementSet(setId)

	if err != nil {
		JsonErrorResponse(w, "couldn't delete a set", err, http.StatusInternalServerError)
		return
	}
}
