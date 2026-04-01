package middlewares

import(
	"strings"
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
		if len(tokenString) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(
				resp.ErrorResponse(false, nil, "Authorization is required"),
			)
		}

		token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
			return configs.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired token",
			})
		}

		// Save user info into context (optional)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("auth_user", claims)
		}

		return c.Next()
	}
}
