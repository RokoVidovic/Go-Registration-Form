package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

type PersonalInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type AddressInfo struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
}

type PaymentsInfo struct {
	Street string `json:"account_owner"`
	IBAN   string `json:"iban"`
}

func handleStep1(c *gin.Context) {
	var data PersonalInfo
	if err := c.BindJSON(&data); err != nil {
		return
	}
	// input verification
	fmt.Println(data)
	c.IndentedJSON(http.StatusCreated, data)
}

func main() {
	var data PersonalInfo
	fmt.Println(data)
	router := gin.Default()
	router.POST("/step1", handleStep1)

	router.Run("localhost:8080")
}
