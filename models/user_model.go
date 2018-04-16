package models

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
