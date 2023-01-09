package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

type SessionRequest struct {
	IdToken string `json:"idToken"`
}

func (s *Router) sessionLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var sessReq SessionRequest

	err := json.NewDecoder(r.Body).Decode(&sessReq)
	if err != nil {
		JsonErrorResponse(w, "couldn't decode session req body", err, http.StatusBadRequest)
		return
	}

	expiresIn := time.Hour * 24 * 7

	cookie, err := s.authClient.GetSessionCookie(r, sessReq.IdToken, expiresIn)
	if err != nil {
		JsonErrorResponse(w, "couldn't get session cookie", err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)
}
