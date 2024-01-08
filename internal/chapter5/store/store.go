package store

import (
	"context"

	"github.com/google/uuid"
	coffeeco "github.com/rubengomes8/learning-ddd-coffeeco/internal/chapter5"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeeco.Product
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (float32, error) {
	dis, err := s.repo.GetStoreDiscount(ctx, storeID)
	if err != nil {
		return 0, err
	}
	return float32(dis), nil
}