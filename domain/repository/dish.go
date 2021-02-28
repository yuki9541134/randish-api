package repository

import (
	"github.com/yuki9541134/randish-api/domain/model"
)

type DishRepository interface {
	GetAll() ([]*model.Dish, error)
}
