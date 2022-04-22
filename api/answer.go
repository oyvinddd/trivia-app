package api

import (
	tapi "github.com/oyvinddd/trivia-api"
	"github.com/oyvinddd/trivia-api/config"
	"net/http"
)

// SubmitAnswer submits an answer for a given question to the API
func SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	cfg, _ := config.New()
	triviaAPI := tapi.New(r.Context(), cfg)
	res, err := triviaAPI.SubmitAnswer(r.Context(), r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}
