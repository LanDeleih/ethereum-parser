package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	restResponse "github.com/landeleih/ethereum-parser/common/types/rest"
)

func (h *Handlers) GetTransactionsHandler(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		restResponse.ToInvalidParamsBadRequest(responseWriter, err)
		return
	}

	var req getTransactionsRequest
	if err = json.Unmarshal(
		requestBody,
		&req,
	); err != nil {
		restResponse.ToInvalidParamsBadRequest(responseWriter, err)
		return
	}

	transactions := h.
		usecase.
		GetTransactions(req.Address)
	if len(transactions) == 0 {
		restResponse.ToBusinessError(responseWriter, &GetTransactionsResponseError{Err: ErrTransactionsNotFound})
		return
	}

	restResponse.ResponseEntity(responseWriter, transactions)
}

type getTransactionsRequest struct {
	Address string `json:"address"`
}

type GetTransactionsResponseError struct {
	Err error `json:"error"`
}

func (r *GetTransactionsResponseError) Error() string {
	return r.Err.Error()
}
