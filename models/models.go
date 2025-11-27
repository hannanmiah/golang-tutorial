package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	Role      string `gorm:"default:user" json:"role"`
	Products  []Product `gorm:"foreignKey:OwnerID" json:"products,omitempty"`
	Orders    []Order `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	Carts     []Cart `gorm:"foreignKey:UserID" json:"carts,omitempty"`
}

type Product struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"default:0" json:"stock"`
	OwnerID     uint    `gorm:"not null" json:"owner_id"`
	Owner       User    `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	CartItems   []Cart  `gorm:"foreignKey:ProductID" json:"cart_items,omitempty"`
	OrderItems  []OrderItem `gorm:"foreignKey:ProductID" json:"order_items,omitempty"`
}

type Order struct {
	gorm.Model
	UserID     uint        `gorm:"not null" json:"user_id"`
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Status     string      `gorm:"default:pending" json:"status"`
	Total      float64     `gorm:"not null" json:"total"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}

type Cart struct {
	gorm.Model
	UserID    uint    `gorm:"not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int     `gorm:"default:1" json:"quantity"`
}