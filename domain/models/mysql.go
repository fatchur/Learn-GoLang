package models

import (
	"database/sql"
	"log"
	"time"
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

func CreateMysqlObj() *Products {
	return &Products{}
}

func (mysqlObj *Products) GetResult(rows *sql.Rows, flag *string) {
	if *flag == "product" {
		for rows.Next() {
			err := rows.Scan(&mysqlObj.Description)
			if err != nil {
				log.Fatal(err)
			}
		}
		// be careful deferring Queries if you are using transactions
		defer rows.Close()
	}
}
