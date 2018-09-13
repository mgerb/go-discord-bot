package model

var Migrations []interface{} = []interface{}{
	&Message{},
	&Attachment{},
	&User{},
	&VideoArchive{},
	&Sound{},
}
