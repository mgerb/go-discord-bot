package db

import (
	"github.com/jinzhu/gorm"

	// database driver for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB

// Init - initialize database
func Init(migrations ...interface{}) {
	var err error
	conn, err = gorm.Open("sqlite3", "data.db")

	if err != nil {
		panic("failed to connect database")
	}
	conn.DB().SetMaxIdleConns(1)
	conn.AutoMigrate(migrations...)
}

// GetConn - get db connection
func GetConn() *gorm.DB {
	return conn
}
