package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Source = "root:123456@tcp(127.0.0.1:3306)/y_traffic?parseTime=True"
)

func NewDail(dataSource string) gorm.Dialector {
	return mysql.Open(dataSource)
}

func NewDb(dial gorm.Dialector) (*gorm.DB, error) {
	return gorm.Open(dial, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
