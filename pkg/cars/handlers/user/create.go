package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//GET THE EMAIL & PASSWORD FROM REQUEST
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}
	//HASH THE PASSWORD
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fauled to hash password",
		})
		return
	}

	//CREATE THE USER
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
        })
		return
	}

	//RESPOND
	c.JSON(http.StatusCreated, gin.H{
        "message": "Successfully created user",
    })
}