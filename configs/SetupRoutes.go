package configs

import (
	"iseng/controller"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(db *mongo.Database) {
	router := gin.Default()
	router.GET("/statement/openingSession", func(c * gin.Context) {controller.GetDailyStatements(c, db.Collection("Statement"))})
	router.Run("localhost:8080")
}