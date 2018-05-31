package db

import (
	"github.com/jinzhu/gorm"

	// database driver for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Conn - database connection
var Conn *gorm.DB

// Init - initialize database
func Init() {
	var err error
	Conn, err = gorm.Open("sqlite3", "data.db")

	if err != nil {
		panic("failed to connect database")
	}

	Conn.DB().SetMaxIdleConns(1)
}
