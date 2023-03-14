package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"rudderstack/internal/api/v1/models"
	"rudderstack/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

func NewDatabaseConnection(cfg *config.Config) (*Database, error) {
	// Set up the database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		"",
	)
	log.Printf("Database connection string: %+v", dsn)

	// Connect to the server to check if database exists
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Check if the database exists, create it if not
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.Database.Name)).Error; err != nil {
		return nil, err
	}

	// Reconnect to the database
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Check if the table exists, create it if not
	migrator := db.Migrator()
	if !migrator.HasTable(&models.TrackingPlan{}) {
		if err := runInitialSQLScript(db); err != nil {
			return nil, err
		}
	}

	// Return the database object
	return &Database{Connection: db}, nil
}

func runInitialSQLScript(db *gorm.DB) error {
	// Read the initial SQL script
	bytes, err := ioutil.ReadFile("./internal/db/initial.sql")
	if err != nil {
		return err
	}

	// Split the SQL script into individual statements
	sqlStatements := strings.Split(string(bytes), ";")

	// Execute each SQL statement
	for _, statement := range sqlStatements {
		trimmed := strings.TrimSpace(statement)
		if trimmed == "" {
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		_, err = sqlDB.Exec(trimmed)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Database) Close() error {
	sqlDb, err := d.Connection.DB()
	if err != nil {
		return err
	}

	return sqlDb.Close()
}
