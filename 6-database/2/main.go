package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Product{})

	if err != nil {
		panic(err)
	}

	// create one
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// create many
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 300.00},
	// 	{Name: "Keyboard", Price: 500.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 3)
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// limit and offset
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 500).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// like
	// var products []Product
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 5)
	// p.Name = "New mouse"
	// db.Save(&p)

	var p2 Product
	db.First(&p2, 5)
	// p2.Name = "New mouse"
	db.Delete(&p2)
}
