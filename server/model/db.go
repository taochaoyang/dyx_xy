package model

import (
	"dyx_xy/server/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := common.TryGetConfig("", "test:209_test@tcp(39.105.96.155:3306)/dyx_xy_test")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true}})
	common.OnError(err, "")
	common.OnError(migrate(DB), "")
}

func migrate(tx *gorm.DB) error {
	return tx.Debug().AutoMigrate(
		&Customer{},
		&User{})
}
