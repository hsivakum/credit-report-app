package main

import (
	"fmt"
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

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
