package store

import (
	"github.com/google/uuid"
	coffeeco "github.com/rubengomes8/learning-ddd-coffeeco/internal"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeeco.Product
}
