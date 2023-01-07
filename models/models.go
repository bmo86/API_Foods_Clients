package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type FoodAllData struct {
	Id          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Price       float64                `json:"price"`
	Ingredients []string               `json:"ingretiens"`
	CreatedAt   *timestamppb.Timestamp `json:"created_at"`
	UpdatedAt   *timestamppb.Timestamp `json:"update_at"`
	Status      bool                   `json:"status"`
}

type Food struct {
	gorm.Model
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status bool    `json:"status"`
}

type FoodWithIngredients struct {
	gorm.Model
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Ingredients []string `json:"ingretiens"`
	Status      bool     `json:"status"`
}

type Ingredients struct {
	gorm.Model
	Name   string `json:"name"`
	Id     int64  `json:"id_food"`
	Status bool   `json:"status"`
}

type FoodUpdate struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type InResIngredients struct {
	Name string `json:"name"`
}
