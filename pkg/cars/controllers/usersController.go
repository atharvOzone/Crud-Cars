package controllers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Login(c *gin.Context) {
	//Get the email and password from the request
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//Compare sent in password with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 *30).Unix(),
	})

	//Sign and get the complete encoded token as a string using the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	//send token
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24*30, "","", false, true)
	c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
    })
}

func Validate (c *gin.Context) {
	user,_ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"Details of the User": user,
	})	
}