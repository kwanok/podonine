package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm")
	}

	err = db.AutoMigrate(&repository.Product{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.Seat{})
	if err != nil {
		return
	}

	repository.Init()
	productRepository := repository.ProductRepository{Db: repository.Gorm}
	product := productRepository.GetProductById(1)
	fmt.Println(product)
	fmt.Println(product.GetSeats())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
