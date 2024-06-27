package handlers

import (
	"net/http"

	restResponse "github.com/landeleih/ethereum-parser/common/types/rest"
)

func (h *Handlers) GetCurrentBlockHandler(
	responseWriter http.ResponseWriter,
	_ *http.Request,
) {
	currentBlock := h.
		usecase.
		GetCurrentBlock()

	restResponse.ResponseEntity(responseWriter, currentBlock)
}
