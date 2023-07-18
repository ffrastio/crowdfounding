package main

import (
	"bwastartup/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	router.GET("/users", UserHandler)
	router.Run()
}

func UserHandler(ctx *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Koneksi berhasil")

	var users []entity.User

	db.Find(&users)

	ctx.JSON(http.StatusOK, users)
}
