package main

import (
	"database/sql"

	_ "github.com/mutecomm/go-sqlcipher/v4"
)

func main() {
	for _, driver := range sql.Drivers() {
		println(driver)
	}
}
