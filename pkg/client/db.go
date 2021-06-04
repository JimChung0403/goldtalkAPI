package client

import (
    "fmt"
    "goldtalkAPI/conf"
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
    DB *gorm.DB
)


type Model struct {
    ID         int `gorm:"primary_key" json:"id"`
    CreatedOn  int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
    DeletedOn  int `json:"deleted_on"`
}

// Setup initializes the database instance
func Setup(config conf.Database) {
    var err error
    DB, err = gorm.Open(config.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        config.User,
        config.Password,
        config.Host,
        config.Name))

    if err != nil {
        log.Fatalf("models.Setup err: %v", err)
    }

    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
        return config.TablePrefix + defaultTableName
    }
    //db.LogMode(true)
    DB.SingularTable(true)
    DB.DB().SetMaxIdleConns(config.IdleConns)
    DB.DB().SetMaxOpenConns(config.OpenConns)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
    defer DB.Close()
}

