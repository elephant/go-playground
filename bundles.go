package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id    int64
	Price float64
}

type Bundle struct {
	Id          string
	Product1    Product
	Product2    Product
	BundlePrice float64
}

func bundleId(product1 Product, product2 Product) string {
	if product1.Id > product2.Id {
		return fmt.Sprintf("%d.%d", product2.Id, product1.Id)
	}

	return fmt.Sprintf("%d.%d", product1.Id, product2.Id)
}

func makeBundles(products []Product, targetBundlePrice float64) []Bundle {
	set := map[string]bool{}
	bundles := []Bundle{}

	for key, product1 := range products {
		if product1.Price >= targetBundlePrice {
			continue
		}

		for i := key + 1; i < len(products); i = i + 1 {

			product2 := products[i]
			if product2.Price >= targetBundlePrice {
				continue
			}

			id := bundleId(product1, product2)
			_, exists := set[id]
			if exists {
				continue
			}

			sum := product1.Price + product2.Price
			if sum <= targetBundlePrice {
				bundle := Bundle{Id: id,
					Product1:    product1,
					Product2:    product2,
					BundlePrice: sum}
				bundles = append(bundles, bundle)
				set[id] = true
			}
		}
	}

	return bundles
}

func main() {
	productJson := `[
		{"id": 1, "price": 20.99}, 
		{"id": 1, "price": 20.99}, 
		{"id": 2, "price": 9.99}, 
		{"id": 3, "price": 29.99}, 
		{"id": 4, "price": 19.99}, 
		{"id": 5, "price": 0.99}    
	]`

	var products []Product

	json.Unmarshal([]byte(productJson), &products)

	//mt.Printf("Products: %#v", products)

	targetBundlePrice := 30.00
	bundles := makeBundles(products, targetBundlePrice)

	for key, bundle := range bundles {
		fmt.Printf("Bundle %d: %#v\n", key, bundle)
	}
}
