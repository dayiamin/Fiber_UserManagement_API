package user

import (
	"time"
	"github.com/golang-jwt/jwt/v5"

	"os"
)


// jwtSecret holds the secret key used for signing JWT tokens.
// It should be set from an environment variable named "JWT_SECRET".
var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // بهتره بعداً توی env بذاری

// GenerateJWT generates a signed JWT token for the given user.
//
// The token includes the following claims:
// - user_id: ID of the user
// - is_admin: Boolean indicating admin status
// - exp: Expiration time (24 hours from now)
// - iat: Issued-at timestamp
//
// Returns the signed token string or an error
func GenerateJWT(userID uint, is_admin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"is_admin" : is_admin,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // اعتبار ۱ روز
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
