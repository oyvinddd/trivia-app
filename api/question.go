package api

import (
	"encoding/json"
	tapi "github.com/oyvinddd/trivia-api"
	"github.com/oyvinddd/trivia-api/config"
	"net/http"
)

// GetDailyQuestion fetches a daily question from the API
func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	cfg, _ := config.New()
	triviaAPI := tapi.New(r.Context(), cfg)
	question, err := triviaAPI.GetDailyQuestion(r.Context())
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
