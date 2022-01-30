package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Products struct {
	Id          int       `json:"id" orm:"id"`
	Name        string    `json:"name" orm:"name"`
	Description string    `json:"description" orm:"description"`
	Images      string    `json:"images" orm:"images"`
	LogoId      string    `json:"logoId" orm:"logoId"`
	CreatedAt   time.Time `json:"createdAt" orm:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" orm:"updatedAt"`
}

func CreateProducts() *Products {
	return &Products{}
}

func (mysqlObj *Products) GetResult(rows *sql.Rows, flagTable *string) error {
	var id sql.NullInt32
	var name sql.NullString
	var description sql.NullString
	var image sql.NullString
	var logoId sql.NullString
	var createdAt sql.NullTime
	var updatedAt sql.NullTime

	if *flagTable == "product" {
		for rows.Next() {
			fmt.Println("-----")
			err := rows.Scan(&id, &name, &description, &image, &logoId)
			if err != nil {
				log.Println("Error in scanning DB", err)
				return err
			}
			if id.Valid {
				mysqlObj.Id = int(id.Int32)
			}
			if name.Valid {
				mysqlObj.Name = name.String
			}
			if description.Valid {
				mysqlObj.Description = description.String
			}
			if image.Valid {
				mysqlObj.Images = image.String
			}
			if logoId.Valid {
				mysqlObj.LogoId = logoId.String
			}
			if createdAt.Valid {
				mysqlObj.CreatedAt = createdAt.Time
			}
			if updatedAt.Valid {
				mysqlObj.UpdatedAt = updatedAt.Time
			}
		}
		// be careful deferring Queries if you are using transactions
		defer rows.Close()
	}
	return nil
}
