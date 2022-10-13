package db

import (
	"database/sql"
	"fmt"

	"simplr-auth/modules/utils"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	env, err := utils.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("env: ", env.ENV)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_database")
	if err != nil {
		return nil, err
	}

	return db, nil
}
