package main

import (
	"credit-report-service-backend-2/controller"
	"credit-report-service-backend-2/db_helper"
	"credit-report-service-backend-2/repository"
	"credit-report-service-backend-2/service"
	"credit-report-service-backend-2/validators"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("hello!")

	db, err := DB()

	if err != nil {
		log.Println("unnable to connect db ", err)
	}

	router := initRoutes(db)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("ValidSSN", validators.ValidSSN)
		_ = v.RegisterValidation("ValidDOB", validators.ValidDOB)
	}

	err = router.Run()
	if err != nil {
		fmt.Printf("%+v", fmt.Errorf("error in starting server, %+v", err))
	}
}

func DB() (*sql.DB, error) {
	db, err := sql.Open("mysql", db_helper.BuildConnectionString())
	if err != nil {
		log.Println("unable to open db")
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
	return db, err
}

func initRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	loginController := controller.NewLoginController()

	appGroup := router.Group("/app")
	{
		appGroup.POST("/login", loginController.Login)
	}

	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)
	registrationController := controller.NewRegistrationService(registrationService)

	registrationGroup := appGroup.Group("/register")
	{
		registrationGroup.POST("", registrationController.CreateUser)
	}
	newDb := sqlx.NewDb(db, "mysql")
	questionsRepository := repository.NewQuestionsRepository(newDb)
	questionsService := service.NewQuestionsService(questionsRepository)
	questionsController := controller.NewQuestionsController(questionsService)

	questionsGroup := appGroup.Group("/questions")
	{
		questionsGroup.GET("", questionsController.GetQuestions)
		questionsGroup.POST("/submit", questionsController.SaveSurveyResults)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "RESOURCE_NOT_FOUND"})
	})

	return router
}
