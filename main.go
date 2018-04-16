package main
import (
    "gin-api-server-sandbox/controllers"
    "github.com/gin-gonic/gin"
)
// main ...
func main() {
    router := gin.Default()
    router.GET("/:id", func(c *gin.Context) { controllers.GetUser(c)})
    router.Run(":8000")
}
