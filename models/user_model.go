package models
import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "gin-api-server-sandbox/config"
    "fmt"
)

var err error
var db *gorm.DB
// init ...
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
    fmt.Printf("connected!!\n")
}

// User is
type User struct {
    ID       int    `json:id`
    Nickname string `json:nickname`
}

// NewUser ...
func NewUser(id int, nickname string) User {
    return User{
        ID:       id,
        Nickname: nickname,
    }
}
// UserRepository is
type UserRepository struct {

}
// NewUserRepository ...
func NewUserRepository() UserRepository {
    return UserRepository{}
}
// GetByID ...
func (m UserRepository) GetByID(id int) *User {
    user := User{}
    user.ID = id
    db.First(&user)
    return &user
}
