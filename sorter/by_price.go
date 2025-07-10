package sorter

import (
	"sort"

	"github.com/monsur4/product-sorter/model"
)

type ByPrice struct{}

func NewByPrice() Sorter {
	return ByPrice{}
}

func (s ByPrice) Sort(products []model.Product) []model.Product {
	sort.SliceStable(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	return products
}
