package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-hca/go-studies/09-rest-api/models"
	"github.com/pedro-hca/go-studies/09-rest-api/utils"
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

func login(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login sucessfull", "token": token})
}
