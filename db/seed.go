package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "gin-api-server-sandbox/config"
    "gin-api-server-sandbox/models"
)

var err error
var db *gorm.DB
var users [5]models.User

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

func userNewRecords() {
  users[0] = models.User{ID:1 , Nickname: "Jinzhu"}
  users[1] = models.User{ID:2 , Nickname: "Nick"}
  users[2] = models.User{ID:3 , Nickname: "Leon"}
  users[3] = models.User{ID:4 , Nickname: "Azrael"}
  users[4] = models.User{ID:5 , Nickname: "Vivaldy"}
  db.Create(users[0])
  db.Create(users[1])
  db.Create(users[2])
  db.Create(users[3])
  db.Create(users[4])
}


func main(){
  db.Set("gorm:table_options", "ENGINE=InnoDB")
  userNewRecords()
}
