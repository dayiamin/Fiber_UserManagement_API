package user

import (


	"github.com/dayiamin/Fiber_UserManagement_API/internal/database"
	"github.com/dayiamin/Fiber_UserManagement_API/internal/models"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)


// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account with name, email, and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  model.UserRegisterRequest  true  "User register data"
// @Success      201   {object} model.UserUserResponse
// @Failure      400   {object} map[string]string
// @Failure      500   {object} map[string]string
// @Router       /register [post]
func Register(c fiber.Ctx) error {
	var body model.UserRegisterRequest

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validateRegister(body); err != nil {
		// برای دیدن خطاها بهتر، می‌تونی خطاها رو استخراج کنی

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	// بررسی تکراری نبودن ایمیل
	var existing model.User
	if err := database.DataBase.Where("email = ?", body.Email).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already registered",
		})
	}

	// هش کردن پسورد
	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user := model.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashed),
		
	}

	if err := database.DataBase.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	response := model.UserUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		
		
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticates a user and returns JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  model.UserLoginRequest  true  "Login credentials"
// @Success      200  {object} map[string]string
// @Failure      400  {object} map[string]string
// @Failure      401  {object} map[string]string
// @Router       /login [post]
func Login(c fiber.Ctx) error{
	var req model.UserLoginRequest

	if err := c.Bind().Body(&req); err !=nil{
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error":"Invalid Input"})
	}

	if err := validateLogin(req); err!=nil{
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error":err})
	}

	var user model.User
	result := database.DataBase.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return c.Status(401).JSON(fiber.Map{"error": "user not found"})
	}

	// بررسی صحت رمز عبور
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "wrong password"})
	}

	token, err := GenerateJWT(user.ID,user.IsAdmin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in generating jwt"})
	}

	
	return c.JSON(fiber.Map{"token": token})

}


// Profile godoc
// @Summary      Get logged-in user's profile
// @Description  Returns user's data by extracting JWT token
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object} map[string]interface{}
// @Failure      404  {object} map[string]string
// @Router       /profile [get]
func Profile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user model.User
	if err := database.DataBase.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
		"is_admin":user.IsAdmin,
	})
}



// GetAllUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all users (admin only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.User
// @Failure      404  {object} map[string]string
// @Router       /users [get]
func GetAllUsers(c fiber.Ctx) error {
	var users []model.User
	if err := database.DataBase.Find(&users).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Users not found"})
	}

	return c.JSON(users)
}