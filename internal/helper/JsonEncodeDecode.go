package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Encode To Json
func JsonEncode[T any](w http.ResponseWriter, r *http.Request, logger *log.Logger, status int, v T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		err = fmt.Errorf("encode json: %w", err)
		ServerError(w, logger, err)
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
