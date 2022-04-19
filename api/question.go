package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	// OTDBResponse response body we get from Open Trivia DB
	OTDBResponse struct {
		Code    int `json:"response_code"`
		Results []struct {
			Category         string   `json:"category"`
			QuestionType     string   `json:"type"`
			Difficulty       string   `json:"difficulty"`
			Question         string   `json:"question"`
			CorrectAnswer    string   `json:"correct_answer"`
			IncorrectAnswers []string `json:"incorrect_answers"`
		} `json:"results"`
	}
	// Question is our domain specific object we serve our end-users
	Question struct {
		Text string `json:"text"`
	}
)

func questionFromResponseBody(body io.ReadCloser) *Question {
	var response OTDBResponse
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil
	}
	if len(response.Results) > 0 {
		first := response.Results[0]
		return &Question{Text: first.Question}
	}
	return nil
}

func createOpenTriviaDBURL(noOfQuestions int) string {
	return fmt.Sprintf("https://opentdb.com/api.php?amount=%d&type=multiple", noOfQuestions)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func GetDailyQuestion(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(createOpenTriviaDBURL(1))
	if err != nil || res.StatusCode != http.StatusOK {
		respondWithError(w, http.StatusInternalServerError, ":(")
		return
	}
	question := questionFromResponseBody(res.Body)
	respondWithJSON(w, http.StatusOK, question)
}
