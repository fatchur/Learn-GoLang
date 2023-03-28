package get

import (
	"fmt"
	"testing"

	mainModel "github.com/inventory/domain/models"
)

func TestGet(t *testing.T) {
	dbType := "mysql"
	dbUsername := "admin"
	dbPassword := "Makanminum12!"
	dbHost := "database-1.c34lvauhheed.us-east-1.rds.amazonaws.com:3306"
	dbDatabase := "sourcesage"
	flagTable := "product"
	query := `SELECT id, name, description, images, logoId FROM sourcesage.Products Where id=1`

	mysqlObj := mainModel.CreateProducts()
	Get(&dbType, &dbUsername, &dbPassword, &dbHost, &dbDatabase, &query, &flagTable, mysqlObj)
	fmt.Println(*mysqlObj)
}
