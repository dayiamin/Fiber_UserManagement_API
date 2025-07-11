package model



import "time"


// User represents the user model stored in the database.
// Used as a response model in admin endpoints.
//
// swagger:model
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"` // رمز هش شده، مخفی در خروجی JSON
	IsAdmin	bool  `json:"is_admin" gorm:"deafult:false"`
	CreatedAt time.Time
}

// UserRegisterRequest is the payload for user registration.
//
// swagger:model
type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	
}

// UserUserResponse is returned after a successful user registration.
//
// swagger:model
type UserUserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	
}

// UserLoginRequest is the payload for logging in a user.
//
// swagger:model
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}