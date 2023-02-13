package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/markusryoti/werk/internal/types"
)

func (s *Handler) registerWorkoutRoutes() {
	s.router.Route("/workouts", func(r chi.Router) {
		r.Use(s.isAuthenticated)
		r.Get("/", s.getWorkoutsByUser)
		r.Post("/", s.addWorkoutHandler)
		r.Get("/{workoutId}", s.getWorkout)
		r.Post("/{workoutId}/addMovement", s.addMovementToWorkout)
		r.Get("/{workoutId}/workoutMovements", s.getMovementsFromWorkout)
		r.Post("/{workoutId}/workoutMovements/{movementId}", s.addSet)
		r.Delete("/{workoutId}", s.removeWorkout)
	})
}

type AddWorkoutRequest struct {
	WorkoutName string `json:"workoutName"`
}

func (s *Handler) addWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newWorkout AddWorkoutRequest

	err := json.NewDecoder(r.Body).Decode(&newWorkout)
	if err != nil {
		JsonErrorResponse(w, "couldn't decode request body", err, http.StatusBadRequest)
		return
	}

	userId := s.getCurrentUser(r)

	err = s.svc.AddNewWorkout(userId, newWorkout.WorkoutName)
	if err != nil {
		JsonErrorResponse(w, "couldn't add a workout", err, http.StatusInternalServerError)
		return
	}

	JsonMessageResponse(w, "workout added", http.StatusOK)
}

func (s *Handler) getWorkoutsByUser(w http.ResponseWriter, r *http.Request) {
	userId := s.getCurrentUser(r)

	workouts, err := s.svc.GetWorkoutsByUser(userId)
	if err != nil {
		JsonErrorResponse(w, "couldn't get workouts", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, workouts, http.StatusOK)
}

func (s *Handler) getWorkout(w http.ResponseWriter, r *http.Request) {
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

type AddMovementRequest struct {
	MovementName string `json:"movementName"`
}

func (s *Handler) addMovementToWorkout(w http.ResponseWriter, r *http.Request) {
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

	movement, err := s.svc.AddNewMovement(workoutId, types.Movement{
		Name: reqBody.MovementName,
		User: userId,
	})
	if err != nil {
		JsonErrorResponse(w, "couldn't add new movement", err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, movement, http.StatusOK)
}

func (s *Handler) removeWorkout(w http.ResponseWriter, r *http.Request) {
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
		s.logger.Infow("users", "workoutUser", workout.User, "reqUser", s.getCurrentUser(r))
		JsonErrorResponse(w, "not authorized", ErrNotAuthorized, http.StatusUnauthorized)
		return
	}

	err = s.svc.DeleteWorkout(workoutId)
	if err != nil {
		JsonErrorResponse(w, "couldn't delete workout", err, http.StatusInternalServerError)
		return
	}
}
