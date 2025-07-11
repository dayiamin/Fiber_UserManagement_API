package user

import (
	
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

)

// JWTMiddleware verifies the presence and validity of a JWT token.
// 
// This middleware checks the "Authorization" header for a Bearer token,
// parses and validates the JWT, and sets the user ID and admin status in context locals.
//
// Routes using this middleware should declare the following Swagger security scheme:
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
//
// Example header:
// Authorization: Bearer <your_token>
func JWTMiddleware() fiber.Handler {
	
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid Authorization header"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			
			return jwtSecret, nil
		})
		
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token", "msg": err.Error(), "token": tokenStr})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		c.Locals("user_id", uint(claims["user_id"].(float64)))
		c.Locals("is_admin",claims["is_admin"])
		return c.Next()
	}
}


// AdminOnly restricts access to routes only for admin users.
//
// This middleware checks the "is_admin" field in context locals (set by JWTMiddleware)
// and returns 403 Forbidden if the user is not an admin.
func AdminOnly()fiber.Handler{
	return func (c fiber.Ctx) error {
		isadmin , ok :=c.Locals("is_admin").(bool)
		if !isadmin || !ok{
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":"Access denied",
			})
		}
		return c.Next()
	}
}

