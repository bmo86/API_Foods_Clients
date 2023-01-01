package repo

import "foods_API_GRPC/models"

type RepoDatabase interface {
	CretedFood(food *models.Food) error
	GetFood(id int64) (*models.Food, error)
	GetFoods(page int64) ([]*models.Food, error)
	UpdateFood(id int64, food models.Food) (bool, error)
	DeleteFood(id int64) (bool, error)
	InsertIngredients(in []*models.Ingredients) error
}

var repository RepoDatabase

func SetRepo(repo RepoDatabase) {
	repository = repo
}

func CretedFood(food *models.Food) error {
	return repository.CretedFood(food)
}

func GetFood(id int64) (*models.Food, error) {
	return repository.GetFood(id)
}

func GetFoods(page int64) ([]*models.Food, error) {
	return repository.GetFoods(page)
}

func UpdateFood(id int64, food models.Food) (bool, error) {
	return repository.UpdateFood(id, food)
}

func DeleteFood(id int64) (bool, error) {
	return repository.DeleteFood(id)
}

//ingredients

func InsertIngredients(in []*models.Ingredients) error {
	return repository.InsertIngredients(in)
}
