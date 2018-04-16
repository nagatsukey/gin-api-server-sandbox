package controllers
import (
    "gin-api-server-sandbox/models"
    "github.com/gin-gonic/gin"
    "reflect"
    "strconv"
)
// User is
type User struct {
}

// Get ...
func (c User) Get(n int) interface{} {
    repo := models.NewUserRepository()
    user := repo.GetByID(n)
    return user
}

func GetUser(c *gin.Context) {
    // Pramを処理する
    n := c.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
        c.JSON(400, err)
        return
    }
    if id <= 0 {
        c.JSON(400, gin.H{"status": "id should be bigger than 0"})
        return
    }
    // データを処理する
    user := User{}
    result := user.Get(id)
    if result == nil || reflect.ValueOf(result).IsNil() {
        c.JSON(404, gin.H{})
        return
    }
    c.JSON(200, result)
}
