package models

type Customer struct {
	ID     uint
	Name   string
	Email  string
	Orders []Order
}

type Product struct {
	ID          uint
	Name        string
	Description string
	CategoryID  uint
	Category    Category
}

type Order struct {
	ID         uint
	CustomerID uint
	Customer   Customer
	Products   []Product `gorm:"many2many:order_products;"`
}
