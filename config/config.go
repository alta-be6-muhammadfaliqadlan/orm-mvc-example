package config

import (
	"immersive6/models/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/immersive6?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	}
}

func MigrateDBAuto() {
	// auto migrate
	DB.AutoMigrate(&user.User{})
}

func MigrateDBV2() {
	// auto migrate
	// DB.Migrator().AddColumn(&User{}, "Alamat")
	// DB.Migrator().AlterColumn(&User{}, "Alamat")

	// DB.Migrator().MigrateColumn(&User{}, )
	// DB.Exec("ALTER TABLE users ADD ID INT;")

}
