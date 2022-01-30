package create

import "github.com/inventory/repository"

func Create(dbtype *string,
	dbUser *string,
	dbPassword *string,
	dbHost *string,
	dbName *string,
	query *string,
	flag *string) error {

	// do business logic here
	repository.MysqlCreate(dbtype, dbUser, dbPassword, dbHost, dbName, query, flag)
	// do business logic here
	return nil
}
