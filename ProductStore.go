package main

import (
	"fmt"
	"sync"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

type ProductStore struct {
	sync.Mutex

	products map[int]Product
	length   int
}

func New() *ProductStore {
	pr := &ProductStore{}
	pr.products = make(map[int]Product)
	pr.length = 0

	return pr
}

func (pr *ProductStore) GetProductById(id int) (Product, error) {
	pr.Lock()
	defer pr.Unlock()

	product, ok := pr.products[id]
	if !ok {
		return Product{}, fmt.Errorf("product with id = %d not found", id)
	} else {
		return product, nil
	}
}

func (pr *ProductStore) GetAllProducts() []Product {
	pr.Lock()
	defer pr.Unlock()

	products := make([]Product, 0, pr.length)

	for _, pr := range pr.products {
		products = append(products, pr)
	}

	return products
}

func (pr *ProductStore) CreateProduct(name, category string, price float64) int {
	pr.Lock()
	defer pr.Unlock()

	newProduct := Product{
		Id:       pr.length,
		Name:     name,
		Category: category,
		Price:    price,
	}

	pr.products[pr.length] = newProduct
	pr.length++

	return newProduct.Id
}

func (pr *ProductStore) DeleteProduct(id int) error {
	pr.Lock()
	defer pr.Unlock()

	if _, ok := pr.products[id]; !ok {
		return fmt.Errorf("product with id = %d not found", id)
	}

	delete(pr.products, id)

	return nil
}

func (pr *ProductStore) GetProductByName(name string) (Product, error) {
	pr.Lock()
	defer pr.Unlock()

	var prod Product

	for _, product := range pr.products {
		if product.Name == name {
			prod = product
		}
	}

	return prod, nil
}

func (pr *ProductStore) DeleteAllProducts() {
	pr.Lock()
	defer pr.Unlock()

	pr.products = make(map[int]Product)

}
