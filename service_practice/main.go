package main

import (
	"go_sample/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.User{})

	//各層のレイヤーの組み立て
	userRepo := repository.
	userSvc := service.
	userH := handler.

	//ルーティング設定
	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("/:id", )
	}
}
