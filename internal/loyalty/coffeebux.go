package loyalty

import (
	"github.com/google/uuid"
	coffeeco "github.com/rubengomes8/learning-ddd-coffeeco/internal"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/store"
)

type CoffeeBux struct {
	ID                                    uuid.UUID
	store                                 store.Store
	coffeeLover                           coffeeco.CoffeeLover
	FreeDrinksAvailable                   int
	RemainingDrinkPurchasesUntilFreeDrink int
}
