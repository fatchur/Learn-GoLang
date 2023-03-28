package get

import (
	"fmt"

	"github.com/inventory/models"
	"github.com/inventory/repository"
)

func Get(dbType *string,
	dbUsername *string,
	dbPassword *string,
	dbHost *string,
	dbDatabase *string,
	dbQuery *string,
	flag *string,
	mysqlObj models.DbConnInteface) error {

	// do business logic here
	err := repository.MysqlGet(dbType, dbUsername, dbPassword,
		dbHost, dbDatabase,
		dbQuery, mysqlObj, flag)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// do business logic here
	return nil
}
