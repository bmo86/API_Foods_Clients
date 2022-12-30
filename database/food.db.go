package database

import (
	"foods_API_GRPC/models"
)

func (i *instacePostgres) CretedFood(food *models.Food) error {
	err := i.db.Create(&food)
	return err.Error
}

func (i *instacePostgres) GetFood(id int64) (*models.Food, error) {

	rows, err := i.db.Raw("SELECT id, name, price FROM food WHERE id = ? ", id).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var food = models.Food{}

	for rows.Next() {
		if err := rows.Scan(&food.Id, &food.Name, &food.Price); err == nil {
			return &food, nil
		}
	}

	if err := rows.Err(); err != nil {
		return &food, nil
	}

	return &food, nil
}

func (i *instacePostgres) GetFoods() ([]*models.Food, error) {
	rows, err := i.db.Raw("SELECT id, name, price FROM food").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var foods []*models.Food

	for rows.Next() {
		var food models.Food
		if err := rows.Scan(&food.Id, &food.Name, &food.Price); err != nil {
			foods = append(foods, &food)
		}
	}

	if rows.Err() != nil {
		return nil, err
	}

	return foods, nil
}

func (i *instacePostgres) UpdateFood(id int64, food *models.Food) (bool, error) {
	err := i.db.Exec("UPDATE food SET name = ?, price = ? WHERE id = ?", food.Name, food.Price, id)

	if err.Error != nil {
		return false, err.Error
	}

	return true, nil
}

func (i *instacePostgres) DeleteFood(id int64) (bool, error) {
	var f = models.Food{}
	if err := i.db.Delete(&f, id).Error; err != nil {
		return false, nil
	}
	return true, nil
}
