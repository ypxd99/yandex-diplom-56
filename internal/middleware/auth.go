package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/ypxd99/yandex-diplom-56/util"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/user/login" || c.Request.URL.Path == "/api/user/register" {
			c.Next()
			return
		}

		logger := util.GetLogger()
		cfg := util.GetConfig()
		cookieName := cfg.Auth.CookieName
		secretKey := []byte(cfg.Auth.SecretKey)

		tokenString, err := c.Cookie(cookieName)
		if err != nil {
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if tokenString == "" {
			c.Next()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.Next()
			return
		}

		c.Set(cookieName, claims.UserID)
		logger.Infof("User authenticated with ID: %s", claims.UserID.String())
		c.Next()
	}
}

func GenerateToken(userID uuid.UUID, key []byte) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieName := util.GetConfig().Auth.CookieName
		userID, exists := c.Get(cookieName)
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		_, ok := userID.(uuid.UUID)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func GetUserID(c *gin.Context) (uuid.UUID, error) {
	cookieName := util.GetConfig().Auth.CookieName
	userID, exists := c.Get(cookieName)
	if !exists {
		return uuid.Nil, errors.New("user ID not found")
	}
	return userID.(uuid.UUID), nil
}
