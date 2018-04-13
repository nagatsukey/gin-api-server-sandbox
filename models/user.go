package models
import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
)
var db *gorm.DB
// init ...
func init() {
    var err error
    DBMS     := "mysql"
    USER     := "root"
    PASS     := "******"
    PROTOCOL := "********8:3306)"
    DBNAME   := "bookshelf"
    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME

    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err)
    }
    fmt.Printf("connected!!")
    //defer db.Close()
}

// User is
type User struct {
    ID       int    `json:id`
    Username string `json:nickname`
}

// NewUser ...
func NewUser(id int, username string) User {
    return User{
        ID:       id,
        Username: username,
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
    //var user = User{ID: id}
    var user User
    db.Where(User{ID: id}).Find(&user)
    return &user
    /*
    if has {
        return &user
    }
    */
    //return nil
}
