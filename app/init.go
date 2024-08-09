package app

import (
	"fmt"
	"github.com/revel/revel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"DineFlowPro/app/models"
	"DineFlowPro/app/middleware"
)

var DB *gorm.DB

func InitDB() {
	dbHost := revel.Config.StringDefault("db.host", "localhost")
	dbPort := revel.Config.StringDefault("db.port", "5432")
	dbUser := revel.Config.StringDefault("db.user", "dineflowpro")
	dbPassword := revel.Config.StringDefault("db.password", "killa1")
	dbName := revel.Config.StringDefault("db.name", "dineflowpro_db")
	dbSSLMode := revel.Config.StringDefault("db.sslmode", "disable")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dbHost, dbPort, dbUser, dbName, dbSSLMode, dbPassword)

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		revel.AppLog.Errorf("Failed to connect to database: %v", err)
		panic("Failed to connect to database")
	}

	// Auto Migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.Order{}, &models.OrderItem{}, &models.MenuItem{}, &models.Menu{}, &models.Delivery{}, &models.Payment{})
	if err != nil {
		revel.AppLog.Errorf("Failed to auto-migrate database schema: %v", err)
		panic("Failed to auto-migrate database schema")
	}

	// Test the connection
	var result int
	if err := DB.Raw("SELECT 1").Scan(&result).Error; err != nil {
		revel.AppLog.Errorf("Failed to execute test query: %v", err)
		panic("Failed to execute test query")
	}

	revel.AppLog.Info("Database connected, schema migrated, and test query executed successfully")
}

func init() {
	revel.OnAppStart(InitDB)

	// Register the authentication interceptor
	revel.InterceptFunc(middleware.AuthenticateUser, revel.BEFORE, "App")
	revel.InterceptFunc(middleware.AuthenticateUser, revel.BEFORE, "OrderTaking")
	revel.InterceptFunc(middleware.AuthenticateUser, revel.BEFORE, "POS")
	revel.InterceptFunc(middleware.AuthenticateUser, revel.BEFORE, "DeliveryConsole")

	// Register role-based authorization interceptors
	revel.InterceptFunc(middleware.AuthorizeRole("admin"), revel.BEFORE, "App")
	revel.InterceptFunc(middleware.AuthorizeRole("staff"), revel.BEFORE, "OrderTaking")
	revel.InterceptFunc(middleware.AuthorizeRole("staff"), revel.BEFORE, "POS")
	revel.InterceptFunc(middleware.AuthorizeRole("driver"), revel.BEFORE, "DeliveryConsole")
}
