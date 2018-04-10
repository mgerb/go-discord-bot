package db

import (
	"github.com/jinzhu/gorm"
	"github.com/mgerb/go-discord-bot/server/config"

	// database driver for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Conn - database connection
var Conn *gorm.DB

// Init - initialize database
func Init() {
	var err error
	Conn, err = gorm.Open("sqlite3", config.Config.Database)

	if err != nil {
		panic("failed to connect database")
	}

	Conn.DB().SetMaxIdleConns(1)
}
