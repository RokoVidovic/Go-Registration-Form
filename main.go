package main

import (
	"Day_2_Go_Registration_Form/core/application/service"
	"Day_2_Go_Registration_Form/infrastructure/persistence/postgres"
	"Day_2_Go_Registration_Form/presentation"
	"context"
	"github.com/gin-gonic/gin"
)

func main() {

	conn := postgres.InitPostgres()
	defer conn.Close(context.Background())

	repo := &postgres.PostgresRegistrationRepo{DB: conn}
	svc := &service.RegistrationService{Repo: repo}
	handler := &presentation.RegistrationController{RegistrationService: svc}

	router := gin.Default()
	router.POST("/step1", handler.HandleStep1)
	router.Run("localhost:8080")
}
