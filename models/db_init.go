package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "gin-api-server-sandbox/config"
)

var err error
var db *gorm.DB

func init() {
    c, err := config.ReadDBConfig("/Users/s01076/go/src/gin-api-server-sandbox/config/db_config.json")
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
