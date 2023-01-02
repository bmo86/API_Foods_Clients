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

func (s *Servepb) GetFood(ctx context.Context, req *foodspb.FoodRequest) (*foodspb.FoodResponse, error) {
	food, err := s.repo.GetFood(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &foodspb.FoodResponse{
		Food: &foodspb.Food{
			Id:          0,
			Name:        food.Name,
			Price:       float32(food.Price),
			Ingredients: food.Ingredients,
			CreatedAt:   food.CreatedAt,
			UpdateAt:    food.UpdatedAt,
			Status:      food.Status,
		},
	}, nil
}

func (s *Servepb) GetFoods(ctx context.Context, req *foodspb.GetFoodsRequest) (*foodspb.FoodsResponse, error) {
	foods, err := s.repo.GetFoods(ctx, req.GetPage())
	if err != nil {
		return nil, err
	}

	var resFoods []*foodspb.Food

	for _, food := range foods {
		f := &foodspb.Food{
			Id:          food.Id,
			Name:        food.Name,
			Price:       float32(food.Price),
			Ingredients: food.Ingredients,
			CreatedAt:   food.CreatedAt,
			UpdateAt:    food.UpdatedAt,
			Status:      food.Status,
		}
		resFoods = append(resFoods, f)
	}

	return &foodspb.FoodsResponse{
		Foods: resFoods,
	}, nil
}

func (s *Servepb) UpdateFood(ctx context.Context, req *foodspb.FoodUpdateRequest) (*foodspb.MessageResponse, error) {
	food := models.FoodUpdate{
		Name:     req.GetName(),
		Price:    float64(req.GetPrice()),
		UpdateAt: timestamppb.Now(),
	}
	res, err := s.repo.UpdateFood(ctx, req.GetId(), food)
	if err != nil {
		return &foodspb.MessageResponse{
			Success: res,
		}, nil
	}

	return &foodspb.MessageResponse{
		Success: res,
	}, nil
}

func (s *Servepb) DeleteFood(ctx context.Context, req *foodspb.FoodRequest) (*foodspb.MessageResponse, error) {
	res, err := s.repo.DeleteFood(ctx, req.GetId())
	if err != nil {
		return &foodspb.MessageResponse{
			Success: res,
		}, nil
	}

	return &foodspb.MessageResponse{
		Success: res,
	}, nil
}
