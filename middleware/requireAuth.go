package middleware

import (
	"fmt"
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("In Middleware")
	//return
	//c.Next()

	// Command + . dile mini hoy code
	// . + Tab dile () ata hoy

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("my_secret_key"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration time
		expirationTime, _ := token.Claims.(jwt.MapClaims)["exp"].(float64)
		if float64(time.Now().Unix()) > expirationTime {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		//result := initializers.DB.First(&user, claims["sub"])
		result := initializers.DB.Where("id = ?", claims["sub"]).Find(&user)

		if result.Error != nil {
			fmt.Println("Error:", result.Error)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// new

		type User struct {
			Username  string
			Id        uuid.UUID `gorm:"type:uuid"`
			Email     string
			IsActive  bool
			CreatedAt time.Time
			DeletedAt time.Time
		}

		// end

		c.Set("user", User{
			Username:  user.Username,
			Id:        user.ID,
			Email:     user.Email,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
		})
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
