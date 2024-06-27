package handlers

import "github.com/landeleih/ethereum-parser/usecase"

type Handlers struct {
	usecase usecase.ParserUseCase
}

func NewHandlers(usecase usecase.ParserUseCase) *Handlers {
	return &Handlers{
		usecase: usecase,
	}
}
