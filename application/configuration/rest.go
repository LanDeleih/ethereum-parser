package configuration

import (
	"net/http"

	"github.com/landeleih/ethereum-parser/rest"
	"github.com/landeleih/ethereum-parser/usecase"
)

func newRestHandlers(
	parserUseCase usecase.ParserUseCase,
) http.Handler {
	return rest.NewChiHandlers(parserUseCase)
}
