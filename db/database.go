package db

import (
    "fmt"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"rudderstack/config"

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
        cfg.Database.Name,
    )
    fmt.Printf(dsn)
    
    // Connect to the database
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Return the database object
    return &Database{Connection: db}, nil
}

func (d *Database) Close() error {
    sqlDb, err := d.Connection.DB()
    if err != nil {
        return err
    }

    return sqlDb.Close()
}
