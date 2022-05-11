package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	// @link https://gorm.io/docs/logger.html
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:mysql.root@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	// Batch Insert
	var users = []User{{Name: "batch inster user 1"}, {Name: "batch inster user 2"}, {Name: "batch inster user 3"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID)
	}

	users = []User{{Name: "CreateInBatches_1"}, {Name: "CreateInBatches_2"}, {Name: "CreateInBatches_100"}}
	// batch size 100
	db.CreateInBatches(users, 100)

	for _, user := range users {
		fmt.Println(user.ID)
	}

	// Create From Map
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "Create From Map", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "batch insert from map 1", "Age": 18},
		{"Name": "batch insert from map 2", "Age": 20},
	})
}
