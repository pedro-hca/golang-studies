package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-hca/go-studies/09-rest-api/models"
)

func signup(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	user.ID = 1

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User Created"})
}
