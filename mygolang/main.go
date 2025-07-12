package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// モデル定義：User と Product を例に
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
	Age   int
	// Products []Product `gorm:"foreignKey:OwnerID"` // 関連付けの練習をする場合
}

type Product struct {
	ID      uint `gorm:"primaryKey"`
	Code    string
	Price   int
	OwnerID uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("example.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 2. マイグレーション
	if err := db.AutoMigrate(&User{}, &Product{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 3. シード済みかカウントで確認
	var userCount int64
	db.Model(&User{}).Count(&userCount)
	if userCount == 0 {
		// User テーブルが空ならダミーデータを投入
		users := []User{
			{Name: "Alice", Email: "alice@example.com", Age: 30},
			{Name: "Bob", Email: "bob@example.com", Age: 25},
			{Name: "Carol", Email: "carol@example.com", Age: 35},
		}
		if err := db.Create(&users).Error; err != nil {
			log.Fatalf("failed to seed users: %v", err)
		}

		products := []Product{
			{Code: "P001", Price: 1000, OwnerID: 1},
			{Code: "P002", Price: 2000, OwnerID: 2},
			{Code: "P003", Price: 1500, OwnerID: 1},
		}
		if err := db.Create(&products).Error; err != nil {
			log.Fatalf("failed to seed products: %v", err)
		}

		fmt.Println("✅ ダミーデータを投入しました")
	} else {
		fmt.Println("ℹ️ 既にシード済み（User 件数:", userCount, "件）")
	}

	//全ユーザー取得からの表示
	var allUsers []User
	db.Find(&allUsers)
	fmt.Println("===全ユーザー===")
	for _, u := range allUsers {
		fmt.Printf("ID:%d  Name:%s  Age:%d\n", u.ID, u.Name, u.Age)
	}
	fmt.Printf("===終了===")

	//条件検索
	var older []User
	db.Where("age >= ?", 30).Find(&older)
	fmt.Println("\n=== Age >= 30 ===")
	for _, u := range older {
		fmt.Printf("%s (Age %d)\n", u.Name, u.Age)
	}

	//ソート処理

}
