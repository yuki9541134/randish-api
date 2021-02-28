package handler

import (
	"encoding/json"
	"github.com/yuki9541134/randish-api/usecase"
	"io"
	"net/http"
)

type DishHandler interface {
	HandleGetAll(w http.ResponseWriter, r *http.Request)
}

type dishHandler struct {
	dishUseCase usecase.DishUseCase
}

func NewDishHandler(useCase usecase.DishUseCase) DishHandler {
	return dishHandler{useCase}
}

func (dh dishHandler) HandleGetAll (w http.ResponseWriter, r *http.Request) {
	products, err := dh.dishUseCase.GetAll()
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(res))
}
