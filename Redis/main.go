// You can edit this code!
// Click here and start typing.
package main

import (
	"paramet/handler"
	"paramet/repositories"
	"paramet/service"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	db := initDatabase()
	redis := initRedis()

	productRepo := repositories.NewproductRepositoryDB(db)
	productService := service.NewcatalogRedis(productRepo , redis)
	prodductHandler := handler.NewcatalogHandler(productService)

	
	app := fiber.New()

	app.Get("/GetProduct" , prodductHandler.GetProduct)

	app.Listen(":8000")
}

func initDatabase() *gorm.DB { 
	dial := mysql.Open("root:mypassword@tcp(localhost:3306)/mydatabase")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "65064384",
	})
}
