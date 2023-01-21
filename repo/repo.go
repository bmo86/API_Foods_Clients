package repo

import (
	"context"
	"foods_API_GRPC/models"
)

type RepoDatabase interface {
	CretedFood(ctx context.Context, food *models.FoodWithIngredients) error
	GetFood(ctx context.Context, id int64) (*models.ResFoodWithIngredients, error)
	GetFoods(ctx context.Context, page int64) ([]*models.ResFoodWithIngredients, error)
	UpdateFood(ctx context.Context, id int64, food models.FoodUpdate) (bool, error)
	DeleteFood(ctx context.Context, id int64) (bool, error)
	InsertIngredients(ctx context.Context, in []*models.Ingredients) error
	GetIngredients(id int64) *models.InResIngredients
	UpdateIngredient(in *models.Ingredients) (bool, error)
	DeleteIngredient(id int64) (bool, error)
}

var repository RepoDatabase

func SetRepo(repo RepoDatabase) {
	repository = repo
}

func CretedFood(ctx context.Context, food *models.FoodWithIngredients) error {
	return repository.CretedFood(ctx, food)
}

func GetFood(ctx context.Context, id int64) (*models.ResFoodWithIngredients, error) {
	return repository.GetFood(ctx, id)
}

func GetFoods(ctx context.Context, page int64) ([]*models.ResFoodWithIngredients, error) {
	return repository.GetFoods(ctx, page)
}

func UpdateFood(ctx context.Context, id int64, food models.FoodUpdate) (bool, error) {
	return repository.UpdateFood(ctx, id, food)
}

func DeleteFood(ctx context.Context, id int64) (bool, error) {
	return repository.DeleteFood(ctx, id)
}

//ingredients

func InsertIngredients(ctx context.Context, in []*models.Ingredients) error {
	return repository.InsertIngredients(ctx, in)
}

func GetIngredients(id int64) *models.InResIngredients {
	return repository.GetIngredients(id)
}

func DeleteIngredient(id int64) (bool, error) {
	return repository.DeleteIngredient(id)
}

func UpdateIngredient(in *models.Ingredients) (bool, error) {
	return repository.UpdateIngredient(in)
}
