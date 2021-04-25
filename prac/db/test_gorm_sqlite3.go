package main

import (
	"gorm.io/gorm"
	"gorm.io/gorm/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D43", Price: 300})

	var product Product

	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	// Update....
	db.Model(&product).Update("Price", 200)
	db.Model(&product).Update(Product{Price: 200, Code: "F42"})

	db.Delete(&product, 1)

}
