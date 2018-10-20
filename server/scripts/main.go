package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Print("Error - Invalid Args\n\nPossible Args:\nrestore-message\nupdate-db\n\n")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "update-db":
		updateDB()
	case "restore-messages":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a channel id")
			return
		}
		restoreMessages(os.Args[2])
	}
}
