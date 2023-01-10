package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	ErrBadRequest = errors.New("bad request")
)

type AddWorkoutRequest struct {
	WorkoutName string
}

func (s *Router) addWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newWorkout AddWorkoutRequest

	err := json.NewDecoder(r.Body).Decode(&newWorkout)
	if err != nil {
		JsonErrorResponse(w, "couldn't decode request body", err, http.StatusBadRequest)
		return
	}

	err = s.repo.AddNewWorkout(newWorkout.WorkoutName)
	if err != nil {
		JsonErrorResponse(w, "couldn't add a workout", err, http.StatusInternalServerError)
		return
	}

	JsonMessageResponse(w, "workout added", http.StatusOK)
}

func (s *Router) getAllWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := s.repo.GetAllWorkouts()
	if err != nil {
		JsonErrorResponse(w, "couldn't get workouts", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, workouts, http.StatusOK)
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

	err = s.repo.AddNewMovement(workoutId, "new movement")
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