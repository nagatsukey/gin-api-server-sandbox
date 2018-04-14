package models
import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
)

var err error
var db *gorm.DB
// init ...
func init() {
    DBMS     := "mysql"
    USER     := "root"
    PASS     := "*******"
    PROTOCOL := "tcp(localhost:3306)"
    DBNAME   := "bookshelf"
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
    user.ID = 1
    db.First(&user)
    fmt.Printf("%s\n", user.Nickname)
    return &user
}
