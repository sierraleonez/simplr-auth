package query

import (
	"database/sql"
	"fmt"

	"simplr-auth/modules/db"

	_ "github.com/go-sql-driver/mysql"
)

func QueryRow(af func(dbArg *sql.DB) (res interface{}, err error)) (res interface{}, err error) {
	db, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	res, err = af(db)
	if err != nil {
		return
	}
	return

	// defer rows.Close()

	// for rows.Next() {
	// 	var user = User{}
	// 	var err = rows.Scan(&user.id, &user.firstName, &user.lastName)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}

	// 	result = append(result, user)
	// }

	// if err = rows.Err(); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
}
