package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
    db.AutoMigrate(&product{})
    mockData(db)
    return productRepositoryRedis{db , redisClient}
}

func (r productRepositoryRedis) GetProduct() (pro []product, err error) {

	key := "repository::GetProducts"

	// Redis Get
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &pro)
		if err == nil {
			fmt.Println("redis")
			return pro, nil
		}
	}

	// Database
	err = r.db.Order("quantity desc").Limit(30).Find(&pro).Error
	if err != nil {
		return nil, err
	}
	
	// Redis Set
	data, err := json.Marshal(pro)
	if err != nil {
		return nil, err
	}

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
	return pro, nil
}