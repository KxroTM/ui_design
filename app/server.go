package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image_Url   string `json:"image_url"`
}

func ProductPage(c *gin.Context) {
	idStart, err := strconv.Atoi(c.DefaultQuery("id_start", "1")) // id_start est un paramètre URL
	if err != nil || idStart == 0 {
		idStart = 1
	} else {
		idStart *= 10 + 3
	}

	var limit int

	if idStart == 1 {
		limit = 12
	} else {
		limit = 12
	}

	products := GetProductsByIdRange(idStart, limit)

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func GetProductsByIdRange(idStart, limit int) []Product {
	rows, err := Db.Query("SELECT * FROM products WHERE id >= ? AND id < ? LIMIT ?", idStart, idStart+limit, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Image_Url)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products
}

func GetTotalProducts() int {
	var count int
	err := Db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count
}
