package main

import (
	"fmt"

	"github.com/mgerb/go-discord-bot/server/db"
)

// keep change script embedded in go files for ease of use
const changeScript = `

-- script to update all users default permission to a value of 1 (user)

UPDATE users SET permissions = 1 WHERE permissions IS NULL OR permissions = 0;

`

func updateDB() {
	db.Init()
	conn := db.GetConn()
	err := conn.Exec(changeScript).Error
	if err != nil {
		fmt.Println(err)
	}
}
