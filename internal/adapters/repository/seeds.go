package repository

import (
	"fmt"
	"time"

	"github.com/RomanshkVolkov/server-storage/internal/core/domain"
	"gorm.io/gorm"
)

func PrintSeedAction(nameTable string, action string) {
	fmt.Println("Seeding table: " + nameTable + " " + action + " Success")
}

func AutoMigrateTable(db *gorm.DB, table interface{}) {
	fmt.Println("AutoMigrateTable")
	isInitialized := db.Migrator().HasTable(&table)
	if !isInitialized {
		db.AutoMigrate(table)
		PrintSeedAction("Shifts", "Create")
	}
}

func RunSeeds(db *gorm.DB) {
	startTimePoint := time.Now().UTC()
	fmt.Println("====================================================================================")
	fmt.Println("Operation run on database", db.Name())
	fmt.Println("Start operation RunSeeds Seeding tables...")
	// db.Exec("DROP TABLE IF EXISTS users_has_kitchens")
	// db.Exec("DROP TABLE IF EXISTS shifts")
	// db.Exec("DROP TABLE IF EXISTS kitchens")
	// db.Exec("DROP TABLE IF EXISTS users")
	// db.Exec("DROP TABLE IF EXISTS user_profiles")
	// db.Exec("DROP TABLE IF EXISTS permissions")
	// db.Exec("DROP TABLE IF EXISTS devs")
	// db.Exec("DROP TABLE IF EXISTS hosting_centers")
	// db.Exec("DROP TABLE IF EXISTS detail_document_tables")
	// db.Exec("DROP TABLE IF EXISTS document_tables")
	// db.Exec("DROP TABLE IF EXISTS detail_documents")
	// db.Exec("DROP TABLE IF EXISTS documents")

	MigrateProcedures(db)

	latency := time.Since(startTimePoint)
	fmt.Println("RunSeeds end operation " + latency.String())
	fmt.Println("====================================================================================")
}

func SeedUsers(db *gorm.DB) {
	AutoMigrateTable(db, &domain.User{})

	var currentRows int64
	db.Model(&domain.User{}).Count(&currentRows)

	if currentRows > 0 {
		return
	}

	users := []domain.User{
		{
			UserData: domain.UserData{
				Username: "dwitmx",
				Email:    "sistemas@dwitmexico.com",
				Name:     "Dwit México",
				IsActive: true,
			},
			Password: GetEnv("ROOT_PASSWORD"),
		},
	}

	for _, user := range users {
		hashedPassword, _ := HashPassword(user.Password)
		user.Password = hashedPassword

		db.Create(&user)
	}
}

func SeedDevAuthorizedIPAddress(db *gorm.DB) {
	AutoMigrateTable(db, &domain.Dev{})

	var currentRows int64
	db.Model(&domain.Dev{}).Count(&currentRows)

	if currentRows > 0 {
		return
	}

	devs := []*domain.Dev{
		{
			IP:  "172.18.0.1",
			Tag: "docker local",
		},
	}

	for _, dev := range devs {
		db.Create(&dev)
	}
}
