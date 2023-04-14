package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

//var rdb *redis.Client

type OrderStatus string

const (
	OrderRequested  OrderStatus = "REQUESTED"
	OrderCreated    OrderStatus = "CREATED"
	OrderInProgress OrderStatus = "IN_PROGRESS"
)

func (s OrderStatus) Value() string {
	return string(s)
}

type Order struct {
	ID              uint64  `gorm:"primaryKey;autoIncrement;column:id;not null"`
	RefCode         string  `gorm:"column:ref_code;not null;uniqueIndex"`
	PolicyID        uint64  `gorm:"column:policy_id;not null"`
	PolicyHolderID  uint64  `gorm:"column:policy_holder_id;not null"`
	InsuredID       *uint64 `gorm:"column:insured_id"`
	InsuredObjectID *uint64 `gorm:"column:insured_object_id"`
	//EffectiveDate         *commons.JSONTime  `gorm:"column:effective_date"`
	//ExpiryDate            *commons.JSONTime  `gorm:"column:expiry_date"`
	PolicyNumber          *string     `gorm:"column:policy_number"`
	ReferencePolicyNumber *string     `gorm:"column:reference_policy_number"`
	Status                OrderStatus `gorm:"column:status;not null"`
	TotalPremium          *float64    `gorm:"column:total_premium;type:double precision"`
	//CreatedAt             commons.JSONTime   `gorm:"column:created_at;autoCreateTime;not null"`
	//UpdatedAt             *commons.JSONTime  `gorm:"column:updated_at"`
	//OrderItems            []*OrderItem
	//SubOrders             []*SubOrder
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) Flatten() map[string]any {
	bytes, _ := json.Marshal(o)
	return map[string]any{
		"id":     o.ID,
		"status": o.Status.Value(),
		"data":   bytes,
	}
}

func Create(ctx context.Context, data *Order) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:33336",
		Password: "pr0d1g1", // no password set
		DB:       0,         // use default DB
	})

	hash := map[string]any{
		"status": data.Status.Value(),
	}
	log.Println(hash)

	fields := data.Flatten()
	key := fmt.Sprintf("orders:%s", data.RefCode)
	_, err := rdb.HSet(ctx, key, fields).Result()
	if err != nil {
		log.Fatalf("error set redis: %v", err)
		return err
	}

	return nil
}

func Retreive(ctx context.Context, key string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:33336",
		Password: "pr0d1g1", // no password set
		DB:       0,         // use default DB
	})

	log.Println(rdb)

	val, err := rdb.HGet(ctx, key, "data").Result()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return err
	}
	fmt.Println("key", val)

	return nil
}
