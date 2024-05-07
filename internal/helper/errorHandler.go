package helper

import (
	"log"
	"net/http"
)

type HttpError struct {
	Message string
}

func ServerError(w http.ResponseWriter, logger *log.Logger, err error) {
	logger.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
