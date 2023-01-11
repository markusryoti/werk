package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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

	movementName := r.URL.Query().Get("name")

	if movementName == "" {
		JsonErrorResponse(w, "empty movement name", ErrBadRequest, http.StatusBadRequest)
		return
	}

	err = s.repo.AddNewMovement(workoutId, movementName)
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

	movements, err := s.repo.GetMovementsFromWorkout(workoutId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get movements", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, movements, http.StatusOK)
}
