package main

import (
	"fmt"

	"github.com/monsur4/product-sorter/service"
	"github.com/monsur4/product-sorter/sorter"
	"github.com/monsur4/product-sorter/utils"
)

func main() {
	products := utils.LoadSampleProducts()

	sorterImpl := sorter.NewByPrice()

	sortService := service.NewSortService(sorterImpl)

	sorted := sortService.Sort(products)
	for _, p := range sorted {
		fmt.Printf("%+v\n", p)
	}
}
