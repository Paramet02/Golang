package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"paramet/repositories"

	"github.com/go-redis/redis/v8"
)

// catalogRedis struct
// โครงสร้างที่ใช้จัดการข้อมูลสินค้า โดยจะพยายามดึงจาก Redis ก่อน ถ้าไม่มีถึงจะไปดึงจากฐานข้อมูล
type catalogRedis struct {
	repo  repositories.ProductRepository // ใช้ดึงข้อมูลจากฐานข้อมูล
	redis *redis.Client                  // ใช้ดึงและบันทึกข้อมูลใน Redis
}

// NewcatalogRedis function
// ฟังก์ชันสำหรับสร้าง catalogRedis โดยรับ repository (repo) และ redis client (redis) มาใช้
func NewcatalogRedis(repo repositories.ProductRepository, redis *redis.Client) CatalogService {
	// คืนค่า catalogRedis ที่พร้อมทำงานตาม CatalogService interface
	return catalogRedis{repo, redis}
}

func (c catalogRedis) GetProduct() (pro []Product, err error) {

	key := "service::GetProducts"

	// Redis Get
	// ดึงข้อมูลจาก Redis ด้วย key "service::GetProducts"
	// หากไม่มีข้อมูลใน Redis (หรือยังไม่เคยตั้งค่า key นี้มาก่อน) จะเกิด error ทำให้ข้ามไปยังขั้นตอนถัดไป
	productsJson, err := c.redis.Get(context.Background(), key).Result()
	if err == nil {
		// Unmarshal แปลง JSON (ในรูปแบบ string) ไปเป็น slice ของ struct Product
		err = json.Unmarshal([]byte(productsJson), &pro)
		if err == nil {
			fmt.Println("redis")
			return pro, nil
		}
	}

	// repository 
	// ดึงข้อมูลจาก Repository (ฐานข้อมูล) และแปลงเป็น struct Product ทีละรายการโดยใช้ for loop
	productDB, err := c.repo.GetProduct()
	if err != nil {
		return nil, err
	}
	for _, p := range productDB {
		pro = append(pro, Product{
			ID: p.ID,
			Name: p.Name,
			Quantity: p.Quantity,
		})
	}

	// Marshal แปลง slice ของ struct Product เป็น JSON string
	data, err := json.Marshal(pro)
	if err != nil {
		return nil, err
	}

	// Set ข้อมูลใน Redis โดยเก็บใน key "service::GetProducts" และกำหนดให้หมดอายุใน 10 วินาที
	err = c.redis.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
	return pro, nil
}