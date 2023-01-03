package database

import (
	"context"
	"foods_API_GRPC/models"

	"gorm.io/gorm"
)

func (i *instacePostgres) CretedFood(ctx context.Context, food *models.FoodWithIngredients) error {
	f := models.Food{
		Model: gorm.Model{
			CreatedAt: food.CreatedAt,
			UpdatedAt: food.UpdatedAt,
		},
		Name:   food.Name,
		Price:  food.Price,
		Status: food.Status,
	}

	err := i.db.Create(&f)

	if err.Error == nil {
		var ingredients []*models.Ingredients

		for _, ingredient := range food.Ingredients {
			i := models.Ingredients{
				Name: ingredient,
				Id:   int64(f.ID),
			}

			ingredients = append(ingredients, &i)
		}

		err := i.InsertIngredients(ctx, ingredients)
		return err
	}

	return err.Error
}

func (i *instacePostgres) GetFood(ctx context.Context, id int64) (*models.FoodWithIngredients, error) {

	sql := "SELECT f.id, f.name, f.price, STRING_AGG(i.name, ', ') AS ingredients FROM foods f INNER JOIN ingredients i ON (i.id = f.id) WHERE f.id = ? GROUP BY f.id"

	rows, err := i.db.Raw(sql, id).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var food = models.FoodWithIngredients{}

	for rows.Next() {
		if err := rows.Scan(&food.ID, &food.Name, &food.Price, &food.Ingredients); err == nil {
			return &food, nil
		}
	}

	if err := rows.Err(); err != nil {
		return &food, nil
	}

	return &food, nil
}

func (i *instacePostgres) GetFoods(ctx context.Context, page int64) ([]*models.Food, error) {
	rows, err := i.db.Raw("SELECT id, name, price FROM food LIMIT = ? OFFSET = ?", page, page*3).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var foods []*models.Food

	for rows.Next() {
		var food models.Food
		if err := rows.Scan(&food.ID, &food.Name, &food.Price); err != nil {
			foods = append(foods, &food)
		}
	}

	if rows.Err() != nil {
		return nil, err
	}

	return foods, nil
}

func (i *instacePostgres) UpdateFood(ctx context.Context, id int64, food models.FoodUpdate) (bool, error) {
	err := i.db.Exec("UPDATE food SET name = ?, price = ? WHERE id = ?", food.Name, food.Price, id)

	if err.Error != nil {
		return false, err.Error
	}

	return true, nil
}

func (i *instacePostgres) DeleteFood(ctx context.Context, id int64) (bool, error) {
	var f = models.Food{}
	if err := i.db.Delete(&f, id).Error; err != nil {
		return false, nil
	}
	return true, nil
}

//ingredients

func (i *instacePostgres) InsertIngredients(ctx context.Context, in []*models.Ingredients) error {
	err := i.db.Create(&in)
	return err.Error
}
