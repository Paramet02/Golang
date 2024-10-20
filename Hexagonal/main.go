package main

import (
	// "net/http"
	"paramet/handlers"
	"paramet/logs"
	"paramet/repository"
	"paramet/service"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



func main() {

	initConfig()
	db := initDatabase()
	logs.Info("banking service started at port " + viper.GetString("app.port"))

	// db.AutoMigrate(&repository.Customer{})
	customerRepositoryDB := repository.NewCustomerRepositoryDB(db) // return CustomerRepository
	customerService := service.NewCustomerService(customerRepositoryDB)
	Handler := handler.NewCustomerHandler(customerService)

	accountRepositoryDB := repository.NewaccountRpeositoryDB(db) 
	accountService := service.NewaccountService(accountRepositoryDB)
	accountHandler := handler.NewacccountHandler(accountService)

	app := fiber.New()
	app.Get("/customers/:id/accounts" , accountHandler.GetAccount)
	app.Post("/customers/:id/accounts" , accountHandler.NewAccount)
	app.Get("/customers", Handler.GetCustomers)
	app.Get("/customer/:id", Handler.GetCustomer)
	app.Listen(fmt.Sprintf(":%d", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDatabase() *gorm.DB {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"), 
		viper.GetInt("database.port"), 
		viper.GetString("database.user"), 
		viper.GetString("database.password"), 
		viper.GetString("database.dbname"))

	// New logger for detailed SQL logging
	// This code snippet is creating a new logger instance for detailed SQL logging in the GORM database
	// connection setup. Here's a breakdown of what each part does:
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	// connect database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})
	// panic : is a builtin function that we can use to stop the flow if a critical situation arises.
	if err != nil {
		panic("fail to connect database")
	}
	
	fmt.Println("Connect Database successful")
	return db
}