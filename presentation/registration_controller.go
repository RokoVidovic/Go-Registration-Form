package presentation

import (
	"Day_2_Go_Registration_Form/core/application/model"
	"Day_2_Go_Registration_Form/core/application/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegistrationController struct {
	RegistrationService *service.RegistrationService
}

func (rc *RegistrationController) HandleStep1(c *gin.Context) {
	var data model.PersonalInfo
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.RegistrationService.SetUUID(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := rc.RegistrationService.SubmitStep1(c.Request.Context(), data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to DB -  " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Step 1 saved!"})
}

func (rc *RegistrationController) HandleStep2(c *gin.Context) {
	var data model.AddressInfo
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.RegistrationService.SubmitStep2(c.Request.Context(), data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to DB"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Step 2 saved!"})
}

func (rc *RegistrationController) HandleStep3(c *gin.Context) {
	var data model.PaymentInfo
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.RegistrationService.SubmitStep3(c.Request.Context(), data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to DB"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Step 3 saved!"})

	// check that all the fields are entered, validate them once more if needed, and move the entry to the final table
}
