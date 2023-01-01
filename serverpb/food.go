package serverpb

import (
	"context"
	"foods_API_GRPC/models"
	"foods_API_GRPC/proto/foodspb"
	"foods_API_GRPC/repo"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Servepb struct {
	repo repo.RepoDatabase
	foodspb.UnimplementedFoodServiceServer
}

func NewServerpb(repo repo.RepoDatabase) *Servepb {
	return &Servepb{repo: repo}
}

func (s *Servepb) CreateFood(ctx context.Context, req *foodspb.CreatedFoodRequest) (*foodspb.FoodResponse, error) {

	t := timestamppb.Now()

	food := &models.Food{
		Name:        req.GetName(),
		Price:       float64(req.GetPrice()),
		Ingredients: req.GetIngredients(),
		CreatedAt:   t,
		UpdatedAt:   t,
		Status:      true,
	}

	if err := s.repo.CretedFood(ctx, food); err != nil {
		return nil, err
	}

	return &foodspb.FoodResponse{
		Food: &foodspb.Food{
			Id:          0,
			Name:        req.GetName(),
			Price:       req.GetPrice(),
			Ingredients: req.GetIngredients(),
			CreatedAt:   t,
			UpdateAt:    t,
			Status:      true,
		},
	}, nil
}
