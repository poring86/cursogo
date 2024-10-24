package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID           int `gorm:"primarykey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	category2 := Category{Name: "Cozinha"}
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		Categories: []Category{category, category2},
	})

	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// })

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, category.Name, "Serial Number", product.SerialNumber.Number)
		}
	}
}
