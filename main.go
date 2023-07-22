package main

import (
	"bwastartup/entity"
	"bwastartup/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_crowdfounding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	userRepository := repository.NewUserRepository(db)

	user := entity.User{
		Name:       "User Baru",
		Email:      "User Baru ",
		Occupation: "Mahasiswa",
	}

	userRepository.Create(user)
}
