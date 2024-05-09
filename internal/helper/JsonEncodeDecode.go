package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Encode To Json
func JsonEncode[T any](w http.ResponseWriter, r *http.Request, status int, v T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		err = fmt.Errorf("encode json: %w", err)
		log.Printf("%v", err)
		JsonEncodeServerError(w, r)
	}
}

// type T should implement the Validator Interface.
// Decode and Validate json
func JsonDecode[T Validator](r *http.Request) (T, map[string]string, error) {
	var decodedJson T
	if err := json.NewDecoder(r.Body).Decode(&decodedJson); err != nil {
		return decodedJson, nil, fmt.Errorf("decode json: %w", err)
	}

	if problems := decodedJson.Valid(r.Context()); len(problems) > 0 {
		return decodedJson, problems, fmt.Errorf("invalid %T: %d problems", decodedJson, len(problems))
	}

	return decodedJson, nil, nil
}

func JsonEncodeError(w http.ResponseWriter, r *http.Request, status int, err error) {
	var httpErr HttpError

	if errors.As(err, &httpErr) {
		errorDetails := []interface{}{}

		for _, v := range httpErr.Message {
			errorDetails = append(errorDetails, v)
		}

		httpRes := HttpErrorResponse{
			Errors:    errorDetails,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		JsonEncode(w, r, status, httpRes)
		return
	}

	log.Printf("%v", err)
	JsonEncodeServerError(w, r)
}

func JsonEncodeServerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := HttpErrorResponse{
		Errors:    []interface{}{"Internal Server Error"},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Printf("JsonEncodeServerError: %v", err)
	}
}
