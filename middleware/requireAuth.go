package middleware

import (
	"crud-go/initializers"
	"crud-go/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	//GET THE COOKIE OFF REQUEST
	tokenString, err := c.Cookie("Authorization")

	if err!= nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//DECODE/VALIDATE THE TOKEN
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")),nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {


		//CHECK THE EXPIRATION
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//FIND THE USER WITH TOKEN SUB
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
            c.AbortWithStatus(http.StatusUnauthorized)
        }

		//ATTACH TO REQUEST
		c.Set("user", user)

		//CONTINUE
		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}