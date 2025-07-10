package sorter

import (
	"sort"

	"github.com/monsur4/product-sorter/model"
)

type ByViews struct{}

func NewByViews() Sorter {
	return ByViews{}
}

func (ByViews) Sort(products []model.Product) []model.Product {
	sort.SliceStable(products, func(i int, j int) bool {
		return products[i].ViewsCount < products[j].ViewsCount
	})

	return products
}
