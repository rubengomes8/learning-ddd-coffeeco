package purchase

import (
	"time"

	"github.com/google/uuid"
	coffeeco "github.com/rubengomes8/learning-ddd-coffeeco/internal"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/payment"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/store"
)

type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
}
