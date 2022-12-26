package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDail(dataSource string) gorm.Dialector {
	return mysql.Open(dataSource)
}

func NewDb(dial gorm.Dialector) (*gorm.DB, error) {
	return gorm.Open(dial, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
