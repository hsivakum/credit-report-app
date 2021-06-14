package main

import (
	"credit-report-service-backend-2/controller"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("hello!")

	err := DB()

	if err != nil {
		log.Println("unnable to connect db ", err)
	}

	router := initRoutes()
	err = router.Run()
	if err != nil {
		fmt.Printf("%+v", fmt.Errorf("error in starting server, %+v", err))
	}
}

func DB() error {
	db, err := sql.Open("mysql", "root:credit-master-password@/credit-db")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		log.Printf("unable to ping db %v", err)
	}
	return err
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
