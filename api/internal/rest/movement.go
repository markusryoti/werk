package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/markusryoti/werk/internal/types"
)

type AddMovementRequest struct {
	MovementName string `json:"movementName"`
}

func (s *Router) addMovementHandler(w http.ResponseWriter, r *http.Request) {
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

	var reqBody AddMovementRequest

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		JsonErrorResponse(w, "couldn't decode request body", err, http.StatusBadRequest)
		return
	}

	if reqBody.MovementName == "" {
		JsonErrorResponse(w, "empty movement name", ErrBadRequest, http.StatusBadRequest)
		return
	}

	userId := s.getCurrentUser(r)

	err = s.svc.AddNewMovement(workoutId, types.Movement{
		Name: reqBody.MovementName,
		User: userId,
	})
	if err != nil {
		JsonErrorResponse(w, "couldn't add new movement", err, http.StatusInternalServerError)
		return
	}
}

func (s *Router) getMovementsFromWorkout(w http.ResponseWriter, r *http.Request) {
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

func (s *Router) removeMovement(w http.ResponseWriter, r *http.Request) {
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
