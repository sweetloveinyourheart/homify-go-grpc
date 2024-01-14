package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(ctx *gin.Context, err error) {
	// Extract validation errors and return 400 response
	var validationErrors []string
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, err.Field()+" must be "+err.Tag())
	}
	ctx.JSON(400, gin.H{"error": "Validation failed", "details": validationErrors})
}

func CustomDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	layout := "2006-01-02"

	_, err := time.Parse(layout, dateStr)
	return err == nil
}
