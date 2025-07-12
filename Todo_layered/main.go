package main

import (
	"Todo_layered/model"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//マイグレーション
	if err := db.AutoMigrate(&model.Todo{}); err != nil{
		log.Fatal(err)
	}

	// 各層のインスタンスの初期化

	//ルーディング設定
	r := gin.Default()
	h.RegisterRouter(r)
}
