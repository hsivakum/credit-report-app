package main

import (
	"fmt"
	"credit-report-service-backend-2/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("hello!")

	router := initRoutes()
	err := router.Run()
	if err != nil {
		fmt.Printf("%+v", fmt.Errorf("error in starting server, %+v", err))
	}
}

func initRoutes() *gin.Engine {
	router := gin.Default()

	loginController := controller.NewLoginController()

	appGroup := router.Group("/app")
	{
		appGroup.POST("/login", loginController.Login)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "RESOURCE_NOT_FOUND"})
	})

	return router
}
