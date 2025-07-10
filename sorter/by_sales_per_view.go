package sorter

import (
	"sort"

	"github.com/monsur4/product-sorter/model"
)

type BySalesPerView struct{}

func NewBySalesPerView() Sorter {
	return BySalesPerView{}
}

func (s BySalesPerView) Sort(products []model.Product) []model.Product {
	sort.SliceStable(products, func(i, j int) bool {
		ratioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		ratioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
		return ratioI > ratioJ
	})
	return products
}
