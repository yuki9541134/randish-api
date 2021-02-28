package usecase

import (
	"github.com/yuki9541134/randish-api/domain/model"
	"github.com/yuki9541134/randish-api/domain/repository"
)

type DishUseCase interface {
	GetAll() ([]*model.Dish, error)
}

type dishUseCase struct {
	dishRepository repository.DishRepository
}

func NewDishUseCase(dr repository.DishRepository) DishUseCase {
	return dishUseCase{dr}
}

func (du dishUseCase) GetAll() ([]*model.Dish, error) {
	return du.dishRepository.GetAll()
}
