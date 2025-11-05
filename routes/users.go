package routes

import (
	"net/http"
	"practice/restapi/models"
	

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
func login(c *gin.Context){
var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse requst data"})
		return
	}
	err:=user.ValidateCredentails()
	

	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not auth"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
