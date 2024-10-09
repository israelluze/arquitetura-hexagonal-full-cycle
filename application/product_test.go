package application_test

import (
	"testing"

	"github.com/israelluze/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	t.Run("Enable a product with a valid price", func(t *testing.T) {
		product := &application.Product{
			Id:     uuid.NewV4().String(),
			Name:   "Product 1",
			Status: application.DISABLED,
			Price:  10,
		}
		err := product.Enable()
		require.Nil(t, err)
		require.Equal(t, application.ENABLED, product.GetStatus())
	})

	t.Run("Enable a product with an invalid price", func(t *testing.T) {
		product := &application.Product{
			Id:     uuid.NewV4().String(),
			Name:   "Product 1",
			Status: application.DISABLED,
			Price:  0,
		}
		err := product.Enable()
		require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
	})
}

func TestProduct_Disable(t *testing.T) {
	t.Run("Disable a product with a valid price", func(t *testing.T) {
		product := &application.Product{
			Name:   "Product 1",
			Status: application.ENABLED,
			Price:  0,
		}
		err := product.Disable()
		require.Nil(t, err)
		require.Equal(t, application.DISABLED, product.GetStatus())
	})

	t.Run("Disable a product with an invalid price", func(t *testing.T) {
		product := &application.Product{
			Name:   "Product 1",
			Status: application.ENABLED,
			Price:  10,
		}
		err := product.Disable()
		require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
	})
}

func TestProduct_IsValid(t *testing.T) {
	t.Run("A valid product", func(t *testing.T) {
		product := &application.Product{
			Id:     uuid.NewV4().String(),
			Name:   "Product 1",
			Status: application.DISABLED,
			Price:  10,
		}
		isValid, err := product.IsValid()
		require.Nil(t, err)
		require.True(t, isValid)
	})

	t.Run("An invalid product", func(t *testing.T) {
		product := &application.Product{
			Name:   "Product 1",
			Status: "INVALID_STATUS",
			Price:  -10,
		}
		_, err := product.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "the status must be enabled or disabled", err.Error())
	})

	t.Run("An invalid product with price < 0", func(t *testing.T) {
		product := &application.Product{
			Name:   "Product 1",
			Status: application.DISABLED,
			Price:  -10,
		}
		_, err := product.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "the price must be greater or equal zero", err.Error())
	})
}
