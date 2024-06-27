package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	restResponse "github.com/landeleih/ethereum-parser/common/types/rest"
)

func (h *Handlers) SubscribeHandler(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		restResponse.ToInvalidParamsBadRequest(responseWriter, err)
		return
	}
	var req subscribeRequest
	if err = json.Unmarshal(
		requestBody,
		&req,
	); err != nil {
		restResponse.ToInvalidParamsBadRequest(responseWriter, err)
		return
	}
	subscribed := h.
		usecase.
		Subscribe(req.Address)
	if !subscribed {
		restResponse.ToBusinessError(responseWriter, ErrSubscriptionFailure)
		return
	}
	restResponse.Created(responseWriter)
}

type subscribeRequest struct {
	Address string `json:"address"`
}
