package main

import (
	"flag"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lon9/go-echo-api-starter/config"
	"github.com/lon9/go-echo-api-starter/database"
	"github.com/lon9/go-echo-api-starter/server"
)

func main() {

	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init(false)
	defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
