package sorter

import (
	"github.com/monsur4/product-sorter/model"
)

type Sorter interface {
	Sort([]model.Product) []model.Product
}

// testing
