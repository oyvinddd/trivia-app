package api

import (
	tapi "github.com/oyvinddd/trivia-api"
	"net/http"
)

func SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	triviaAPI := tapi.New(r.Context())
	res, err := triviaAPI.SubmitAnswer(r.Context(), r.Body)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, res)
	}
}
