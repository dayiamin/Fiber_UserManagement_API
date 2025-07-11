package user

import (
	"github.com/dayiamin/Fiber_UserManagement_API/internal/models"
	"github.com/go-playground/validator/v10"
)

// validate is an instance of the validator used for struct validation
var validate = validator.New()


// validateRegister validates the fields of a UserRegisterRequest struct.
//
// Returns a map of field names to error messages if validation fails,
// or nil if the input is valid.
//
// Validations:
// - Email: must be a valid email format
// - Password: must have at least 6 characters
// - Name: cannot be empty
func validateRegister(body model.UserRegisterRequest) map[string]string {
	
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		
		errors := make(map[string]string)
		for _, e := range validationErrors {
			switch e.Field() {
			case "Email":
				errors["email"] = "Email must be Valid"
			case "Password":
				errors["password"] = "password must contains at least 6 charecters"
			case "Name":
				errors["name"] = "name can not be empty"
			default:
				errors[e.Field()] = "invalid input"
			}
		}
		return errors
	}

	return nil
}


// validateLogin validates the fields of a UserLoginRequest struct.
//
// Returns a map of field names to error messages if validation fails,
// or nil if the input is valid.
//
// Validations:
// - Email: must be a valid email format
// - Password: must have at least 6 characters
func validateLogin(body model.UserLoginRequest) map[string]string {
	
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		
		errors := make(map[string]string)
		for _, e := range validationErrors {
			switch e.Field() {
			case "Email":
				errors["email"] = "Email must be Valid"
			case "Password":
				errors["password"] = "password must contains at least 6 charecters"
		
			default:
				errors[e.Field()] = "invalid input"
			}
		}
		return errors
	}

	return nil
}
