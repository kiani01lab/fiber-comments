package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kiani01lab/fiber-comments/config"
)

func HandleJWT(c *fiber.Ctx) error {
	key, ok := c.GetReqHeaders()["X-Api-Key"]
	if !ok {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "missing or malformed jwt"})
	}

	claims, err := parseKey(key[0])
	if err != nil {
		return err
	}

	if isKeyExpired(claims) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "invalid token"})
	}

	return c.Next()
}

func isKeyExpired(claims jwt.MapClaims) bool {
	return time.Now().Unix() > int64(claims["exp"].(float64))
}

func parseKey(key string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unauthorized")
		}
		return []byte(config.Config("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	if !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	return claims, nil
}
