package statementController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getDailyStatements(c *gin.Context) {
	c.IndentedJSON(http.StatusOK)
}