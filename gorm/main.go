package main

import (
	"context"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CoverageTypeEnum uint8
type MultiplierTypeEnum uint8

const (
	Tsi MultiplierTypeEnum = iota + 1
	Up
)

type TestTable struct {
	gorm.Model
	Code  string
	Price uint
}

type SubTable struct {
	gorm.Model
	MainTableID   uint64           `gorm:"column:main_table_id;not null"`
	CoverageType  CoverageTypeEnum `gorm:"column:coverage_type;type:int;not null"`
	Name          string           `gorm:"column:name;type:varchar(100);not null"`
	Description   *string          `gorm:"column:description;type:text"`
	MinSumInsured float64          `gorm:"column:min_sum_insured;type:double precision;default:0;not null"`
	MaxSumInsured *float64         `gorm:"column:max_sum_insured;type:double precision"`
	CreatedBy     uint64           `gorm:"column:created_by;not null"`
	UpdatedBy     *uint64          `gorm:"column:updated_by"`
	DeletedBy     *uint64          `gorm:"column:deleted_by"`

	//ForeignTable ForeignTable `gorm:"foreignKey:ForeignTableID"`
}

type MainTable struct {
	gorm.Model
	CpBenefitID    uint64             `gorm:"column:cp_benefit_id;not null"`
	Name           string             `gorm:"column:name;type:varchar(100);not null"`
	Short          string             `gorm:"column:short;type:varchar(15)"`
	Description    *string            `gorm:"column:description;type:text"`
	IsPrimary      bool               `gorm:"column:is_primary;type:boolean;not null"`
	MultiplierType MultiplierTypeEnum `gorm:"column:multiplier_type;type:int;not null"`
	CreatedBy      uint64             `gorm:"column:created_by;not null"`
	UpdatedBy      *uint64            `gorm:"column:updated_by"`
	DeletedBy      *uint64            `gorm:"column:deleted_by"`
	SubTable       []*SubTable
}

func main() {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	dsn := "host=localhost user=prodigi password=pr0d1g1 dbname=pws-mv-pricing-service port=11116 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()
	mainTable := &MainTable{}
	/*
		if err := db.WithContext(ctx).Where(&MainTable{Name: "dd"}).First(mainTable).Error; err != nil {
			log.Fatal(err)
		}*/
	db.WithContext(ctx).Preload(clause.Associations).Where(&MainTable{Name: "dd"}).Find(&mainTable).First(mainTable)

	log.Println("------------------", mainTable)
	log.Println("=============", mainTable.SubTable[0])

	/*
		ft := ForeignTable{}
		ft.ID = 1
		db.Model(&ft).Updates(ForeignTable{Name: "nomer1", CreatedBy: 2})
	*/
	/*
		desc := "test"
		Max := 2.2
		var by uint64 = 1

		subTable := SubTable{
			MainTableID:   1,
			CoverageType:  1,
			Name:          "test",
			Description:   &desc,
			MinSumInsured: 0.9,
			MaxSumInsured: &Max,
			CreatedBy:     1,
			UpdatedBy:     &by,
			DeletedBy:     nil,
		}
		var subTableList []*SubTable
		subTableList = append(subTableList, &subTable)

		mainTable := MainTable{
			CpBenefitID:    2,
			Name:           "dd",
			Short:          "dd",
			Description:    &desc,
			IsPrimary:      true,
			MultiplierType: Tsi,
			CreatedBy:      by,
			UpdatedBy:      &by,
			DeletedBy:      &by,
			SubTable:       subTableList,
		}

		// Migrate the schema
		db.AutoMigrate(&MainTable{})
		db.AutoMigrate(&SubTable{})

		//db.Create(&ft)
		db.Create(&mainTable)
	*/
	/*
		//db.AutoMigrate(&tmp)
		db.Create(&TestTable{Code: "D43", Price: 1000})

		var product TestTable
		db.First(&product, 1)

		var product2 TestTable
		db.First(&product2, "code = ?", "D44")

		log.Println(product2)
	*/

	// Create
	/*
		db.Create(&Product{Code: "D42", Price: 100})

		// Read
		var product Product
		db.First(&product, 1)                 // find product with integer primary key
		db.First(&product, "code = ?", "D42") // find product with code D42

		// Update - update product's price to 200
		db.Model(&product).Update("Price", 200)
		// Update - update multiple fields
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
		db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

		// Delete - delete product
		db.Delete(&product, 1)
	*/
}
