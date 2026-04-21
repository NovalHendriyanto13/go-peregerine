package middlewares

import(
	"strings"
	"crypto/subtle"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
	"peregerine/configs"
)

func JwtProtected() fiber.Handler{
	return func(c *fiber.Ctx) error {
		var resp responses.BaseResponse

		authHeader := c.Get("Authorization")
		
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				resp.ErrorResponse(false, nil, "Authorization is required"),
			)
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				resp.ErrorResponse(false, nil, "Authorization must be Bearer token"),
			)
		}

		token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
			// FIX: Validate signing method — prevents alg:none attacks
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return configs.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(
				resp.ErrorResponse(false, nil, "Invalid or expired token"),
			)
		}

		// Save user info into context (optional)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("auth_user", claims)
		}

		return c.Next()
	}
}

func AppKey() fiber.Handler {
	return func(c *fiber.Ctx) error{
		var resp responses.BaseResponse

		appKeyHeader := c.Get("x-app-key")
		
		if appKeyHeader == "" {
			return c.Status(fiber.StatusBadRequest).JSON(
				resp.ErrorResponse(false, nil, "Invalid request header"),
			)
		}

		if subtle.ConstantTimeCompare([]byte(appKeyHeader), []byte(configs.XAppKey)) != 1 {
			return c.Status(fiber.StatusBadRequest).JSON(
				resp.ErrorResponse(false, nil, "Invalid request header"),
			)
		}

		return c.Next()
	}
}
