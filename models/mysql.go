package models

import "database/sql"

// DbQuery ...
type DbConnInteface interface {
	GetResult(*sql.Rows, *string) error
}

/*
type DbConn struct {
	DbType           string `json:"dbType"`
	DbConnectionAddr string `json:"dbConnectionAddr"`
	DbObject         *sql.DB
}

func CreateDbConn(dbType *string, username *string, password *string, host *string) *DbConn {
	myDbConn := DbConn{}
	myDbConn.DbType = *dbType
	myDbConn.DbConnectionAddr = *username + ":" + *password + "@" + *host
	return &myDbConn
}

func (dbConn *DbConn) DoQuery(query *string) error {
	db, err := sql.Open(dbConn.DbType, dbConn.DbConnectionAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()
	result, err := db.Query(*query)
	if err != nil {
		panic(err.Error())
		return err
	}

	repository.RowsToStructs(result)
	// be careful deferring Queries if you are using transactions
	defer result.Close()

	return nil
}
*/
