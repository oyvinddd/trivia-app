package api

import (
	"encoding/json"
	tapi "github.com/oyvinddd/trivia-api"
	"github.com/oyvinddd/trivia-api/config"
	"net/http"
	"os"
)

// GetDailyQuestion fetches a daily question from the API
func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	cfg := config.Firebase(
		os.Getenv("FB_TYPE"),
		os.Getenv("FB_PROJECT_ID"),
		os.Getenv("FB_PRIVATE_KEY_ID"),
		os.Getenv("FB_PRIVATE_KEY"),
		os.Getenv("FB_CLIENT_EMAIL"),
		os.Getenv("FB_CLIENT_ID"))
	question, err := tapi.New(r.Context(), cfg).GetDailyQuestion()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, question)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&payload)
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}
