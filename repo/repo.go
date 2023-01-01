package repo

import (
	"context"
	"foods_API_GRPC/models"
)

type RepoDatabase interface {
	CretedFood(ctx context.Context, food *models.Food) error
	GetFood(ctx context.Context, id int64) (*models.Food, error)
	GetFoods(ctx context.Context, page int64) ([]*models.Food, error)
	UpdateFood(ctx context.Context, id int64, food models.Food) (bool, error)
	DeleteFood(ctx context.Context, id int64) (bool, error)
	InsertIngredients(ctx context.Context, in []*models.Ingredients) error
}

var repository RepoDatabase

func SetRepo(repo RepoDatabase) {
	repository = repo
}

func CretedFood(ctx context.Context, food *models.Food) error {
	return repository.CretedFood(ctx, food)
}

func GetFood(ctx context.Context, id int64) (*models.Food, error) {
	return repository.GetFood(ctx, id)
}

func GetFoods(ctx context.Context, page int64) ([]*models.Food, error) {
	return repository.GetFoods(ctx, page)
}

func UpdateFood(ctx context.Context, id int64, food models.Food) (bool, error) {
	return repository.UpdateFood(ctx, id, food)
}

func DeleteFood(ctx context.Context, id int64) (bool, error) {
	return repository.DeleteFood(ctx, id)
}

//ingredients

func InsertIngredients(ctx context.Context, in []*models.Ingredients) error {
	return repository.InsertIngredients(ctx, in)
}
