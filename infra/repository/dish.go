package repository

import (
	"database/sql"
	"github.com/yuki9541134/randish-api/domain/model"
	"github.com/yuki9541134/randish-api/domain/repository"
)

type dishRepository struct {
	db *sql.DB
}

func NewDishRepository(db *sql.DB) repository.DishRepository {
	return dishRepository{db}
}

func (pr dishRepository) GetAll() ([]*model.Dish, error) {
	query := `select id, name, description from dishes limit 10;`
	rows, err := pr.db.Query(query)
	if err != nil {
		panic(err)
	}

	var dishes []*model.Dish
	for rows.Next() {
		var (dish model.Dish)
		err := rows.Scan(&dish.ID, &dish.Name, &dish.Description)
		if err != nil {
			panic(err)
		}
		dishes = append(dishes, &dish)
	}

	return dishes, nil
}
