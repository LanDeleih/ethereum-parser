package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

func ToInvalidParamsBadRequest(
	responseWriter http.ResponseWriter,
	message error,
) {
	responseWriter.WriteHeader(http.StatusBadRequest)
	responseWriter.Header().Set(contentTypeHeader, contentTypeJSON)
	response, marshalErr := json.Marshal(message)
	if marshalErr != nil {
		log.Println("can't marshal response error", marshalErr)
	}
	if _, err := responseWriter.Write(response); err != nil {
		log.Println("can't response back", err)
	}
}

func ToBusinessError(
	responseWriter http.ResponseWriter,
	message error,
) {
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Set(contentTypeHeader, contentTypeJSON)
	response, marshalErr := json.Marshal(message.Error())
	if marshalErr != nil {
		log.Println("can't marshal response error", marshalErr)
	}
	if _, err := responseWriter.Write(response); err != nil {
		log.Println("can't response back", err)
	}
}

func Created(
	responseWriter http.ResponseWriter,
) {
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Header().Set(contentTypeHeader, contentTypeJSON)
	if _, err := responseWriter.Write(nil); err != nil {
		log.Println("can't response back", err)
	}
}

func ResponseEntity(
	responseWriter http.ResponseWriter,
	body any,
) {
	response, err := json.Marshal(body)
	if err != nil {
		ToBusinessError(responseWriter, err)
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Set(contentTypeHeader, contentTypeJSON)
	if _, err = responseWriter.Write(response); err != nil {
		log.Println("can't response back", err)
	}
}
