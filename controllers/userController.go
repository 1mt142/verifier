package controllers

import (
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Signup(c *gin.Context) {

	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{Email: body.Email, Password: string(hash), Username: body.Username}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "User created",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "email=?", body.Email)
	//println("user ", user)
	//if user.ID == 0 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Invalid email or password",
	//	})
	//	return
	//}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Stop
	//token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
	//	"sub": user.ID,
	//	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	//})
	//println("Token:::", token)
	//tokenString, err := token.SignedString([]byte("SDSFDFEFERGRBT"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Token fail",
	//	})
	//	return
	//}
	//

	// Define the claims for the JWT
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Email,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	// Create a new token object with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Set the signing key
	key := []byte("my_secret_key")
	tokenString, err := token.SignedString(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't generate token!",
		})
	}
	// *** token based authentication
	//c.JSON(http.StatusOK, gin.H{
	//	"access_token": tokenString,
	//})

	// *** cookie based authentication
	c.SetSameSite(http.SameSiteLaxMode)
	// name, value string, maxAge int, path, domain string, secure, httpOnly bool
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged in!",
		"data":    user,
	})

}
