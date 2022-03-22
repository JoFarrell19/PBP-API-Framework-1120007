package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {

	db := Connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM product")
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	var product Product
	var products []Product

	for results.Next() {
		err = results.Scan(&product.Code, &product.Name, &product.Qty)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	if len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}
func AddProduct(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var product Product

	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return
	}
	db.Exec("INSERT INTO product (code, name, qty) VALUES (?,?,?)", product.Code, product.Name, product.Qty)

	c.IndentedJSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var product Product

	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return
	}

	result, errQuery := db.Exec("UPDATE product SET name=?, qty=? WHERE code=?", product.Name, product.Qty, product.Code)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Failed to Update")
			return
		} else {
			c.IndentedJSON(http.StatusOK, product)
		}
	}
}

func DeleteProduct(c *gin.Context) {
	db := Connect()
	defer db.Close()

	code := c.Query("Code")

	result, errQuery := db.Exec("DELETE FROM product WHERE Code=?", code)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Failed to Delete")
		} else {
			c.IndentedJSON(http.StatusOK, code)
		}
	}
}
