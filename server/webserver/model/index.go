package model

// Migrations - list of database migrations
var Migrations = []interface{}{
	&Message{},
	&Attachment{},
	&User{},
	&VideoArchive{},
	&Sound{},
	&UserEventLog{},
}
