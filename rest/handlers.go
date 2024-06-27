package rest

import (
	"net/http"

	parserhandlers "github.com/landeleih/ethereum-parser/rest/parser"
	"github.com/landeleih/ethereum-parser/usecase"
)

func NewChiHandlers(
	parserUseCase usecase.ParserUseCase,
) http.Handler {
	mux := http.NewServeMux()

	handlers := parserhandlers.NewHandlers(
		parserUseCase,
	)

	mux.HandleFunc("/apis/v1/block/current", handlers.GetCurrentBlockHandler)
	mux.HandleFunc("/apis/v1/transaction", handlers.GetTransactionsHandler)
	mux.HandleFunc("/apis/v1/subscribe", handlers.SubscribeHandler)

	return mux
}
