package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "gin-api-server-sandbox/config"
    "gin-api-server-sandbox/models"
)

var err error
var db *gorm.DB

func init() {
    c, err := config.ReadDBConfig("./config/db_config.json")
    if err != nil {
        panic(err)
    }
    DBMS     := c.DBMS
    USER     := c.User
    PASS     := c.Pass
    PROTOCOL := c.Protcol
    DBNAME   := c.DBName
    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
    db, err = gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err)
    }
}

func main(){
    db.Set("gorm:table_options", "ENGINE=InnoDB")
    db.AutoMigrate(&models.User{})
}
