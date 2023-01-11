package database

import (
	"context"
	"foods_API_GRPC/models"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetIngredients(t *testing.T) {
	c := require.New(t)

	re, err := NewConnectionDatabase("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		c.Error(err, "connection db")
	}
	expected := models.InResIngredients{
		Id:   []string{"1", "2", "3"},
		Name: []string{"test ingredient", "test ingredient", "test ingredient3"},
	}

	result := re.GetIngredients(1)
	if err != nil {
		c.Error(err, expected)
	}
	c.Equal(expected.Id, result.Id)
	c.Equal(expected.Name, result.Name)
}

func TestUpdateIngredient(t *testing.T) {
	c := require.New(t)
	i, err := NewConnectionDatabase("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		c.Error(err, "connection db")
	}

	in := &models.Ingredients{

		Model: gorm.Model{ID: 3, UpdatedAt: time.Now()},
		Name:  "test ingredient3",
	}

	// Call the UpdateIngredient function with the test data
	updated, err := i.UpdateIngredient(in)
	if err != nil {
		t.Errorf("Error updating ingredient: %v", err)
	}

	// Check that the function returns true
	if !updated {
		t.Errorf("Expected function to return true but got false")
	}
}

func TestDelete(t *testing.T) {
	c := require.New(t)
	i, err := NewConnectionDatabase("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		c.Error(err, "connection db")
	}

	delete, err := i.DeleteFood(context.Background(), 28)
	if err != nil {
		t.Log(err.Error())
	}
	c.Equal(true, delete)
}
