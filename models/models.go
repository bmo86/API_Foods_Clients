package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Id          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Price       float64                `json:"price"`
	Ingredients []string               `json:"ingretiens"`
	CreatedAt   *timestamppb.Timestamp `json:"created_at"`
	UpdatedAt   *timestamppb.Timestamp `json:"update_at"`
	Status      bool                   `json:"status"`
}

type FoodWithoutIngredients struct {
	gorm.Model
	Name      string                 `json:"name"`
	Price     float64                `json:"price"`
	CreatedAt *timestamppb.Timestamp `json:"created_at"`
	UpdateAt  *timestamppb.Timestamp `json:"update_at"`
	Status    bool                   `json:"status"`
}

type Ingredients struct {
	gorm.Model
	Name   string `json:"name"`
	IDFood int64  `json:"id_food"`
	Status bool   `json:"status"`
}

type FoodUpdate struct {
	gorm.Model
	Name     string                 `json:"name"`
	Price    float64                `json:"price"`
	UpdateAt *timestamppb.Timestamp `json:"update_at"`
}
