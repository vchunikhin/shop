package shop

import (
	"github.com/google/uuid"
	"math/big"
	"time"
)

type BaseEntity struct {
	Id      uuid.UUID `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted time.Time `json:"deleted"`
}

type User struct {
	BaseEntity
	FullName string `json:"full_name"`
}

type Category struct {
	BaseEntity
	Name string `json:"name"`
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
