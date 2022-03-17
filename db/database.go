package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Realiza la conexión
var dsn = "root:root@tcp(localhost:3306)/users_go?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error en la conexion")
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}

}()