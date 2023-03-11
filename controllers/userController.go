package controllers

import (
	"fmt"
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
	"github.com/1mt142/verifier/services"
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
	// Generate OTP
	otpIs := services.GenerateOTP()
	// Send OTP :: TODO stop this service for unusual email sending
	// err = services.SendOTPViaEmail(body.Email, otpIs, "OTP for user verification in verifier app")
	// if err != nil {
	//	fmt.Println(err)
	//  }
	// Store OTP
	otp := models.OTP{Otp: otpIs, Channels: "Email", SenderId: user.ID}
	otpResult := initializers.DB.Create(&otp)
	if otpResult.Error != nil {
		fmt.Println("OTP Store Fail")
	}
	//
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
	result := initializers.DB.First(&user, "email=?", body.Email)
	// check if user was created successfully
	if result.Error != nil {
		// handle error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
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
		"sub":   user.ID,
		"name":  user.Username,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
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
	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
	})
}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged in!",
		"data":    user,
	})

}
