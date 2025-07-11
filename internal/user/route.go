package user

import "github.com/gofiber/fiber/v3"



// RegisterRoutes sets up the API routes for the application.
//
// Routes:
// - POST   /register → Register a new user (public)
// - POST   /login    → Login and receive a JWT (public)
//
// Protected routes (require valid JWT):
// - GET    /profile  → Get current user profile
//
// Admin-only routes (require JWT and admin privilege):
// - GET    /admin/users → Get list of all users
//
// Middleware:
// - JWTMiddleware: Validates JWT token and extracts user data
// - AdminOnly: Ensures the user has admin privileges
func RegisterRoutes(router fiber.Router) {
	router.Post("/register", Register)
	router.Post("/login", Login)
	protected := router.Group("", JWTMiddleware())
	protected.Get("/profile", Profile)
	admin := router.Group("/admin", JWTMiddleware(),AdminOnly())
	admin.Get("/users", GetAllUsers)
}