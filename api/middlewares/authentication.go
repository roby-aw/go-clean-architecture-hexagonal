package middlewares

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/roby-aw/go-clean-architecture-hexagonal/utils"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gofiber/fiber/v2"
)

func MiddleJWTUser(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "invalid token",
		})
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	secret := os.Getenv("JWT_SECRET")
	key, err := utils.Decode(secret)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "invalid token",
		})
	}

	header, _, err := jose.Decode(tokenString, key)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "invalid token",
		})
	}

	var str utils.JwtTokenClaimsUser
	err = json.Unmarshal([]byte(header), &str)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "invalid token",
		})
	}

	timenow := time.Unix(str.Exp, 0)
	if timenow.Before(time.Now()) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "invalid token",
		})
	}

	if str.Role != "Admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "invalid token",
		})
	}
	c.Locals("id", str.Sub)
	c.Locals("role", str.Role)
	return c.Next()
}
