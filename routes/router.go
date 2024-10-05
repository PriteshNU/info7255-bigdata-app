package routes

import (
	"info7255-bigdata-app/database"
	"info7255-bigdata-app/handlers"
	"info7255-bigdata-app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	redisRepo := database.NewRedisRepository("localhost:6379")
	planService := services.NewPlanService(redisRepo)
	planHandler := handlers.NewPlanHandler(planService)

	v1 := router.Group("/v1")
	{
		v1.POST("/plan", planHandler.CreatePlan)
		v1.GET("/plan/:objectId", planHandler.GetPlan)
		v1.DELETE("/plan/:objectId", planHandler.DeletePlan)
		v1.GET("/plans", planHandler.GetAllPlans)
	}

	return router
}
