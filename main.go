package main

import (
	controller "framework/Controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/getProduct", controller.GetProduct)
	router.POST("/addProduct", controller.AddProduct)
	router.PUT("/updateProduct", controller.UpdateProduct)
	router.DELETE("/deleteProduct", controller.DeleteProduct)
	router.Run()
}
