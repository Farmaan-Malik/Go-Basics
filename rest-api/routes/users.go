package routes

import (
	"example/rest-api/models"
	"example/rest-api/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func userSignup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"status": "failure", "message": fmt.Sprint(err)})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(500, gin.H{"status": "failure", "message": fmt.Sprint(err)})
		return
	}
	ctx.JSON(201, gin.H{"status": "success", "data": user})

}

func userLogin(ctx *gin.Context) {
	var u models.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(400, gin.H{"status": "failure", "message": fmt.Sprint(err)})
	}
	fmt.Println(u.Id)
	err = u.ValidateCredentials()
	if err != nil {
		ctx.JSON(401, gin.H{"status": "failure", "message": fmt.Sprint(err)})
		return
	}
	fmt.Println(u.Id)
	token, err := utils.GenerateJwt(u.Email, u.Id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(401, gin.H{"status": "failure", "message": "could not create token"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Logged in successfully!", "data": token})
}
