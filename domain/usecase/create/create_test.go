package create

import (
	"fmt"
	"testing"
	"time"

	mainModel "github.com/inventory/domain/models"
)

func TestCreate(t *testing.T) {
	dbType := "mysql"
	dbUsername := "admin"
	dbPassword := "Makanminum12!"
	dbHost := "database-1.c34lvauhheed.us-east-1.rds.amazonaws.com:3306"
	dbDatabase := "sourcesage"
	flagTable := "product"

	createdAt := time.Now().UTC().Format("2006-01-02 15:04:05.999999")
	query := fmt.Sprintf(`INSERT into sourcesage.Products 
		(name, description, images, logoId, createdAt, updatedAt) 
		values ("sugar", "sugar related products", "s3://sourcesage/...", "1", "%s", "%s")`, createdAt, createdAt)
	fmt.Println(query)

	mysqlObj := mainModel.CreateProducts()
	Create(&dbType, &dbUsername, &dbPassword, &dbHost, &dbDatabase, &query, &flagTable)
	fmt.Println(*mysqlObj)
}
