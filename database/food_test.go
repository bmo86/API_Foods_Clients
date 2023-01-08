package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetIngredients(t *testing.T) {
	c := require.New(t)

	re, err := NewConnectionDatabase("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		c.Error(err, "connection db")
	}
	expected := []string{"harina", "leche", "azucar"}

	result := re.GetIngredients(1)
	if err != nil {
		c.Error(err, expected)
	}
	t.Log(result)
	c.Equal(expected, result)
}
