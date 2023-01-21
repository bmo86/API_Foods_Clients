package serverpb

import (
	"context"
	"foods_API_GRPC/models"
	"foods_API_GRPC/proto/foodspb"
	"foods_API_GRPC/repo"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Servepb struct {
	repo repo.RepoDatabase
	foodspb.UnimplementedFoodServiceServer
}

func NewServerpb(repo repo.RepoDatabase) *Servepb {
	return &Servepb{repo: repo}
}

func (s *Servepb) CreatedFood(ctx context.Context, req *foodspb.CreatedFoodRequest) (*foodspb.FoodResponse, error) {

	t := time.Now().UTC()

	food := &models.FoodWithIngredients{
		Model: gorm.Model{
			CreatedAt: t,
			UpdatedAt: t,
		},
		Name:        req.GetName(),
		Price:       float64(req.GetPrice()),
		Ingredients: req.GetIngredients(),
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
			CreatedAt:   timestamppb.New(t),
			UpdateAt:    timestamppb.New(t),
			Status:      true,
		},
	}, nil
}

func (s *Servepb) GetFood(ctx context.Context, req *foodspb.FoodRequest) (*foodspb.FoodGetResponse, error) {
	food, err := s.repo.GetFood(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &foodspb.FoodGetResponse{
		Food: &foodspb.ResFood{
			Id:    0,
			Name:  food.Name,
			Price: float32(food.Price),
			Ingredi: &foodspb.Ingre{
				Id:   food.Ingredients.Id,
				Name: food.Ingredients.Name,
			},
			CreatedAt: timestamppb.New(food.CreatedAt),
			UpdateAt:  timestamppb.New(food.UpdatedAt),
			Status:    food.Status,
		},
	}, nil
}

func (s *Servepb) GetFoods(ctx context.Context, req *foodspb.GetFoodsRequest) (*foodspb.FoodsResponse, error) {
	foods, err := s.repo.GetFoods(ctx, req.GetPage())
	if err != nil {
		return nil, err
	}

	var resFoods []*foodspb.ResFood

	for _, food := range foods {
		f := &foodspb.ResFood{
			Id:        int64(food.ID),
			Name:      food.Name,
			Price:     float32(food.Price),
			Ingredi:   &foodspb.Ingre{Id: food.Ingredients.Id, Name: food.Ingredients.Name},
			CreatedAt: timestamppb.New(food.CreatedAt),
			UpdateAt:  timestamppb.New(food.UpdatedAt),
			Status:    food.Status,
		}
		resFoods = append(resFoods, f)
	}

	return &foodspb.FoodsResponse{
		Foods: resFoods,
	}, nil
}

func (s *Servepb) UpdateFood(ctx context.Context, req *foodspb.FoodUpdateRequest) (*foodspb.MessageResponse, error) {
	food := models.FoodUpdate{
		Model: gorm.Model{
			UpdatedAt: time.Now().UTC(),
		},
		Name:  req.GetName(),
		Price: float64(req.GetPrice()),
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
