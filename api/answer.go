package api

import (
	tapi "github.com/oyvinddd/trivia-api"
	"github.com/oyvinddd/trivia-api/config"
	"net/http"
	"os"
)

// SubmitAnswer submits an answer for a given question to the API
func SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	cfg := config.Firebase(
		os.Getenv("FB_TYPE"),
		os.Getenv("FB_PROJECT_ID"),
		os.Getenv("FB_PRIVATE_KEY_ID"),
		os.Getenv("FB_PRIVATE_KEY"),
		os.Getenv("FB_CLIENT_EMAIL"),
		os.Getenv("FB_CLIENT_ID"))
	res, err := tapi.New(r.Context(), cfg).SubmitAnswer(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}
