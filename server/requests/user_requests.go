package requests

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/server/handlers"
	"ComputerWorld_API/server/responses"
	"net/http"
	"regexp"
)

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func ValidateUserInputs(user *models.User) error {
	errVI := isValidUserInput(user)
	if errVI != nil {
		return errVI
	}

	// Validate Password
	errVP := isValidUserPassword(user.Password)
	if errVP != true {
		return responses.NewHTTPError(http.StatusBadRequest,
			"Password is invalid: "+
				"Must contain at least 8 characters, "+
				"including 1 uppercase, "+
				"1 lowercase, "+
				"1 number,"+
				" and 1 special character")
	}

	// After validation, hash the string
	hash, errHash := userPasswordHash(user.Password)
	if errHash != nil {
		return errHash
	}

	// Replace the current password string with the hashed password
	user.Password = hash

	return nil
}

func isValidUserInput(user *models.User) error {
	validEmailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matchedEmail, _ := regexp.MatchString(validEmailPattern, user.Email)
	if !matchedEmail {
		return responses.NewHTTPError(http.StatusBadRequest, "Email is invalid: Incorrect Formatting")
	}
	validNamePattern := `^[a-zA-Z\s]+$`
	matchedName, _ := regexp.MatchString(validNamePattern, user.Name)
	if !matchedName {
		return responses.NewHTTPError(http.StatusBadRequest, "Name is invalid: No Special Characters or Numbers are allowed")
	}

	return nil
}

func isValidUserPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)

	return hasLowercase && hasUppercase && hasDigit && hasSpecial
}

func userPasswordHash(password string) (string, error) {
	salt, err := handlers.CreateSalt(16)
	if err != nil {
		return "", responses.NewHTTPError(http.StatusBadRequest, "Error Generating Salt")
	}

	hashedPassword := handlers.HashPassword(password, salt)
	return hashedPassword, nil
}
