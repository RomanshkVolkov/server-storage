package repository

import (
	"fmt"

	"gorm.io/gorm"
)

// stored procedures

func ExistSP(db *gorm.DB, nombreSP string) bool {
	var existe int
	err := db.Raw("SELECT COUNT(*) FROM sys.procedures WHERE name = ?", nombreSP).Scan(&existe).Error

	if err != nil {
		fmt.Println("error when verifying the existence of the stored procedure: %w", err)
		return false
	}

	return existe > 0
}

func ExistFunc(db *gorm.DB, nombreFunc string) bool {
	var existe int
	err := db.Raw("SELECT COUNT(*) FROM sys.objects WHERE name = ? AND type = 'FN'", nombreFunc).Scan(&existe).Error

	if err != nil {
		fmt.Println("error when verifying the existence of the function: %w", err)
		return false
	}

	return existe > 0
}

func ExistTable(db *gorm.DB, name string) bool {
	var rowCount int
	err := db.Raw("SELECT COUNT(*) FROM sys.tables WHERE name = ?", name).Scan(&rowCount).Error

	fmt.Println(rowCount)
	if err != nil {
		fmt.Println("error when verifying the existence of the table: %w", err)
		return false
	}

	return rowCount > 0
}

func MigrateProcedures(db *gorm.DB) {

}
