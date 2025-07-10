package service

import (
	"github.com/monsur4/product-sorter/model"
	"github.com/monsur4/product-sorter/sorter"
)

type SortService struct {
	Sorter sorter.Sorter
}

func NewSortService(sorter sorter.Sorter) *SortService {
	return &SortService{Sorter: sorter}
}

func (s *SortService) Sort(products []model.Product) []model.Product {
	return s.Sorter.Sort(products)
}
