package purchase

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	coffeeco "github.com/luispinto23/ddd-coffeeco/internal"
	"github.com/luispinto23/ddd-coffeeco/internal/payment"
	"github.com/luispinto23/ddd-coffeeco/internal/store"
)

type Purchase struct {
	ID                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	total              money.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}
