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
	dbConfig := fmt.Sprintf(
		"root:%s@tcp(%s:%d)/%s",
		env.DB_PASS,
		env.DB_HOST,
		env.DB_PORT,
		env.DB_NAME,
	)
	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
