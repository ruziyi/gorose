package main

import (
	"github.com/gohouse/gorose"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"./config"
)

func main() {
	DB := gorose.Connect.Open(config.Configs,"mysql_dev")
	// close DB
	defer DB.Close()
	// get the db chaining object
	var db gorose.Database

	res := db.Table("users").Count()
	max := db.Table("users").Max("age")
	db.Table("users").Min("age")
	db.Table("users").Avg("age")
	db.Table("users").Sum("age")
	fmt.Println(db.LastSql())
	fmt.Println(res)
	fmt.Println(max)

}

