package controllers

import (
	"fmt"
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
	"github.com/1mt142/verifier/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Signup(c *gin.Context) {
	log.Info().Msg("Test Log")

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
		log.Error().Err(result.Error).Msg("DB__CREATE_ERROR")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	// Generate OTP
	otpIs := services.GenerateOTP()
	// Send OTP :: TODO stop this service for unusual email sending
	err = services.SendOTPViaEmail(body.Email, otpIs, "OTP for user verification in verifier app")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("OTP", otpIs)
	// Store OTP
	otp := models.OTP{Otp: otpIs, Channels: "Email", SenderId: user.ID}
	otpResult := initializers.DB.Create(&otp)
	if otpResult.Error != nil {
		fmt.Println("OTP Store Fail")
	}
	//
	c.Header("UserId", user.ID.String())
	c.Set("otp_users", user.ID)
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
	// check otp verified
	if result != nil && user.IsVerified != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please verify your account",
		})
		return
	}
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

func OtpVerify(c *gin.Context) {

	var otpStruct struct {
		Otp string
	}
	if c.Bind(&otpStruct) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read OTP",
		})
		return
	}

	if len(strings.TrimSpace(otpStruct.Otp)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OTP can not be empty.",
		})
		return
	}

	otpToken := c.GetHeader("Otp-Token") // extract the value of the 'otp_token' header
	var otpValue models.OTP
	initializers.DB.First(&otpValue, "sender_id = ?", otpToken)
	// Check OTP
	if otpValue.Otp == otpStruct.Otp {
		var user models.User
		initializers.DB.Find(&user, "id = ?", otpToken)
		// check if already verified
		if user.IsVerified == true {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Already Verified",
			})
			return
		}
		// update IsVerified Status
		initializers.DB.Model(&user).Updates(models.User{IsVerified: true})
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome! You are now Verified",
		})
		return
	}
	if otpValue.Otp != otpStruct.Otp {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect OTP",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": "Technical Error",
	})
	return

}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged in!",
		"data":    user,
	})

}

func GetUsers(c *gin.Context) {
	var count int64
	limit := c.Query("limit")
	offset := c.Query("offset")
	_limit, err := strconv.Atoi(limit)
	if err != nil {
		// handle error
	}
	//
	_offset, err := strconv.Atoi(offset)
	if err != nil {
		// handle error
	}

	var users []models.User
	initializers.DB.
		Select("id", "created_at", "username", "email", "is_verified").
		Limit(_limit).
		Offset(_offset).
		Find(&users).
		Count(&count)
	//
	type Response struct {
		Id         uuid.UUID
		Username   string
		Email      string
		IsActive   bool
		IsVerified bool
	}

	//var users []models.User

	// Assume that the users slice has been populated with data

	var finalData []Response
	for _, user := range users {
		newData := Response{
			Id:         user.ID,
			Username:   user.Username,
			Email:      user.Email,
			IsActive:   user.IsActive,
			IsVerified: user.IsVerified,
		}
		finalData = append(finalData, newData)
	}
	//
	c.JSON(http.StatusOK, gin.H{
		"count":   count,
		"limit":   _limit,
		"offset":  _offset,
		"data":    finalData,
		"message": "User results found",
	})
}
