package requests

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/server/handlers"
	"ComputerWorld_API/server/responses"
	"net/http"
	"regexp"
)

type AdminRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func ValidateAdminInputs(admin *models.Admin) error {
	errVI := isValidAdminInput(admin)
	if errVI != nil {
		return errVI
	}

	// Validate Password
	errVP := isValidAdminPassword(admin.Password)
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
	hash, errHash := adminPasswordHash(admin.Password)
	if errHash != nil {
		return errHash
	}

	// Replace the current password string with the hashed password
	admin.Password = hash

	return nil
}

func isValidAdminInput(admin *models.Admin) error {
	validEmailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matchedEmail, _ := regexp.MatchString(validEmailPattern, admin.Email)
	if !matchedEmail {
		return responses.NewHTTPError(http.StatusBadRequest, "Email is invalid: Incorrect Formatting")
	}
	validNamePattern := `^[a-zA-Z\s]+$`
	matchedName, _ := regexp.MatchString(validNamePattern, admin.Name)
	if !matchedName {
		return responses.NewHTTPError(http.StatusBadRequest, "Name is invalid: No Special Characters or Numbers are allowed")
	}

	return nil
}

func isValidAdminPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)

	return hasLowercase && hasUppercase && hasDigit && hasSpecial
}

func adminPasswordHash(password string) (string, error) {
	salt, err := handlers.CreateSalt(16)
	if err != nil {
		return "", responses.NewHTTPError(http.StatusBadRequest, "Error Generating Salt")
	}

	hashedPassword := handlers.HashPassword(password, salt)
	return hashedPassword, nil
}
