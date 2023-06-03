package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductServer struct {
	store *ProductStore
}

func NewServer() *ProductServer {
	store := New()

	return &ProductServer{store: store}
}

func main() {
	router := gin.Default()
	server := NewServer()

	router.GET("/products/", server.getAllProductsHandler)
	router.GET("/products/:id", server.getProductById)
	router.POST("/products/create/", server.createProductHandler)
	router.DELETE("/products/delete/:id", server.deleteProductById)

	router.Run(":8080")

}

func (ps *ProductServer) getAllProductsHandler(c *gin.Context) {
	allProducts := ps.store.GetAllProducts()
	c.IndentedJSON(http.StatusOK, allProducts)
}

func (ps *ProductServer) deleteProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "id = %d not found", id)
	}
	delete(ps.store.products, id)
}

func (ps *ProductServer) getProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		fmt.Print(err.Error())
	}

	product, erro := ps.store.GetProductById(id)

	if erro != nil {
		c.String(http.StatusBadRequest, erro.Error())
	}

	c.IndentedJSON(http.StatusOK, product)
}

func (ps *ProductServer) createProductHandler(c *gin.Context) {
	type RequestProd struct {
		Name     string  `json:"name"`
		Category string  `json:"category"`
		Price    float64 `json:"price"`
	}
	var p RequestProd

	if err := c.ShouldBindJSON(&p); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id := ps.store.CreateProduct(p.Name, p.Category, p.Price)

	c.IndentedJSON(http.StatusOK, gin.H{"Id": id})
}
