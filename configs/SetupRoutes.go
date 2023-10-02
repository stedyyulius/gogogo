package configs

import (
	"iseng/controller"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(db *mongo.Database) {
	router := gin.Default()
	router.GET("/statement/openingSession", func(c * gin.Context) {controller.GetDailyStatements(c, db.Collection("Statement"))})
	
	router.GET("/masterbet", func(c * gin.Context) {controller.GetDailyStatements(c, db.Collection("MasterBet"))})
	
	pprof.Register(router, "debug/pprof")
	
	router.Run("localhost:8080")

	
}
