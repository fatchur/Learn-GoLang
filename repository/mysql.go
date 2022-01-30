package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/inventory/models"
)

func MysqlGet(dbtype *string,
	dbUser *string,
	dbPassword *string,
	dbHost *string,
	dbName *string,
	query *string,
	sqlObj models.DbConnInteface,
	flagTable *string) error {

	address := fmt.Sprintf("%s:%s@tcp(%s)/%s", *dbUser, *dbPassword, *dbHost, *dbName)
	db, err := sql.Open(*dbtype, address)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Query(*query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	sqlObj.GetResult(result, flagTable)
	return nil
}

func MysqlCreate(dbtype *string,
	dbUser *string,
	dbPassword *string,
	dbHost *string,
	dbName *string,
	query *string,
	flagTable *string) error {

	address := fmt.Sprintf("%s:%s@tcp(%s)/%s", *dbUser, *dbPassword, *dbHost, *dbName)
	db, err := sql.Open(*dbtype, address)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec(*query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
	return nil
}
