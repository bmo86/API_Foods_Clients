package models

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Ingredients []string  `json:"ingretiens"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
	Status      bool      `json:"status"`
}

type FoodWithoutIngredients struct {
	gorm.Model
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Status    bool      `json:"status"`
}

type Ingredients struct {
	gorm.Model
	Name   string `json:"name"`
	IDFood int64  `json:"id_food"`
	Status bool   `json:"status"`
}
