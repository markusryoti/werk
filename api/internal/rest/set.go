package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/markusryoti/werk/internal/types"
)

type AddSetRequest struct {
	Reps   uint8 `json:"reps"`
	Weight uint8 `json:"weight"`
}

func (s *Router) addSet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody AddSetRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		JsonErrorResponse(w, "invalid message body", err, http.StatusBadRequest)
		return
	}

	movementIdStr := chi.URLParam(r, "movementId")
	movementId, err := strconv.ParseUint(movementIdStr, 10, 64)

	err = s.svc.AddMovementSet(movementId, types.Set{
		Reps:   reqBody.Reps,
		Weight: reqBody.Weight,
	})

	if err != nil {
		JsonErrorResponse(w, "couldn't add a new set", err, http.StatusInternalServerError)
		return
	}
}

func (s *Router) removeSet(w http.ResponseWriter, r *http.Request) {
	setIdStr := chi.URLParam(r, "setId")

	setId, err := strconv.ParseUint(setIdStr, 10, 64)
	if err != nil {
		JsonErrorResponse(w, "invalid delete request", err, http.StatusBadRequest)
		return
	}

	err = s.svc.DeleteMovementSet(setId)

	if err != nil {
		JsonErrorResponse(w, "couldn't delete a set", err, http.StatusInternalServerError)
		return
	}
}
