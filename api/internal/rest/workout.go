package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	ErrBadRequest    = errors.New("bad request")
	ErrNotAuthorized = errors.New("not authorized")
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

	uid, _ := r.Context().Value(userKey).(string)

	err = s.svc.AddNewWorkout(uid, newWorkout.WorkoutName)
	if err != nil {
		JsonErrorResponse(w, "couldn't add a workout", err, http.StatusInternalServerError)
		return
	}

	JsonMessageResponse(w, "workout added", http.StatusOK)
}

func (s *Router) getAllWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := s.svc.GetAllWorkouts()
	if err != nil {
		JsonErrorResponse(w, "couldn't get workouts", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, workouts, http.StatusOK)
}

func (s *Router) getWorkout(w http.ResponseWriter, r *http.Request) {
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

	workout, err := s.svc.GetWorkout(workoutId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get workout", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, workout, http.StatusOK)
}

func (s *Router) removeWorkout(w http.ResponseWriter, r *http.Request) {
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

	workout, err := s.svc.GetWorkout(workoutId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get workout", err, http.StatusInternalServerError)
		return
	}

	if workout.User != s.getCurrentUser(r) {
		JsonErrorResponse(w, "not authorized", ErrNotAuthorized, http.StatusUnauthorized)
		return
	}

	err = s.svc.DeleteWorkout(workoutId)
	if err != nil {
		JsonErrorResponse(w, "couldn't delete workout", err, http.StatusInternalServerError)
		return
	}
}
