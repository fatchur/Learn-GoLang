package repository

import (
	"database/sql"
	"fmt"

	"github.com/inventory/models"
)

func MysqlDoQuery(sqlObj models.DbConnInteface, flag *string) {
	db, err := sql.Open("mysql", "admin:Makanminum12!@tcp(database-1.c34lvauhheed.us-east-1.rds.amazonaws.com:3306)/sourcesage")
	if err != nil {
		fmt.Println(err)
	}

	query := `SELECT name, description, images, logoId FROM sourcesage.Products Where id=1`

	defer db.Close()
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	sqlObj.GetResult(result, flag)

}
