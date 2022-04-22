package api

import (
	"encoding/json"
	tapi "github.com/oyvinddd/trivia-api"
	"github.com/oyvinddd/trivia-api/config"
	"net/http"
	"os"
)

var (
	firebaseConfig = config.Firebase(
		os.Getenv("FB_TYPE"),
		os.Getenv("FB_PROJECT_ID"),
		os.Getenv("FB_PRIVATE_KEY_ID"),
		os.Getenv("FB_PRIVATE_KEY"),
		os.Getenv("FB_CLIENT_EMAIL"),
		os.Getenv("FB_CLIENT_ID"))
)

// TriviaHandler is our main entry point for the Trivia API
func TriviaHandler(w http.ResponseWriter, r *http.Request) {
	triviaAPI := tapi.New(r.Context(), firebaseConfig)
	switch r.Method {
	case http.MethodGet:
		handleGetDailyQuestion(w, triviaAPI)
		break
	case http.MethodPost:
		handleSubmitAnswer(w, r, triviaAPI)
		break
	default:
		respondWithStatus(w, http.StatusMethodNotAllowed)
	}
}

func handleGetDailyQuestion(w http.ResponseWriter, api *tapi.TriviaAPI) {
	question, err := api.GetDailyQuestion()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, question)
}

func handleSubmitAnswer(w http.ResponseWriter, r *http.Request, api *tapi.TriviaAPI) {
	result, err := api.SubmitAnswer(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, result)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&payload)
}

func respondWithStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Write([]byte(nil))
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}
