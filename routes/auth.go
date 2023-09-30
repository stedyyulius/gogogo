package auth
import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine)
auth := route.Group("/"){
    auth.POST("/login")
}