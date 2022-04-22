package shop

import (
	"github.com/google/uuid"
	"math/big"
	"time"
)

type BaseEntity struct {
	Id      uuid.UUID `json:"-"`
	Created time.Time `json:"-"`
	Updated time.Time `json:"-"`
	Deleted time.Time `json:"-"`
}

type User struct {
	//BaseEntity
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}

type Category struct {
	BaseEntity
	Name string `json:"name" binding:"required"`
}

type Order struct {
	BaseEntity
	User        *User     `json:"user"`
	TotalAmount big.Float `json:"total_amount"`
}

type Product struct {
	BaseEntity
	Name     string    `json:"name"`
	Price    big.Float `json:"price"`
	Sku      string    `json:"sku"`
	Category *Category `json:"category"`
}
