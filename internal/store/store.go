package store

import (
	"github.com/google/uuid"
	coffeeco "github.com/luispinto23/ddd-coffeeco/internal"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeeco.Product
}
