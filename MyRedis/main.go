// You can edit this code!
// Click here and start typing.
package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	handlers "github.com/paramet02/Handlers"
	"github.com/paramet02/Repository"
	service "github.com/paramet02/Service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	db := initDatabase()

	productRepo := repository.NewproductRepository(db)
	productSv := service.NewproService(productRepo)
	productHL := handlers.NewproductHandler(productSv)


	app := fiber.New()

	app.Post("/Product" , productHL.CreateProduct)
	app.Patch("/Product/:id", productHL.UpdateProduct)
	app.Get("/Products" , productHL.GetProducts)
	app.Get("/Product" , productHL.GetProduct)

	app.Listen(":8000")
}

func initDatabase() *gorm.DB { 
	// Connection string for PostgreSQL
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	dial := postgres.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "mypassword",
	})
}
