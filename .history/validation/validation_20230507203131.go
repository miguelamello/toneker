package validation

import (
	"encoding/json"
	//"fmt"
	"regexp"
	s "strings"
)

type Payload struct {
	Email   string `json:"email"`
	Token   string `json:"token"`
	Created string `json:"created"`
	Expires string `json:"expires"`
}

// Object for 200 responses
type Response200 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload Payload `json:"payload"`
}

// Object for 400 responses
type Response400 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Verify if the email string is valid
func isEmail(email string) bool {
	// Define a regular expression for validating email addresses
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	// Check if the email address matches the regular expression
	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}

// Verify if the email string exists and is valid
func ValidateEmailString(email string) bool {
	cEmail := s.Trim(email, " ")
	if len(cEmail) == 0 { 
		return false 
	} else { 
		if isEmail(cEmail) {
			return true 
		} else {			
			return false
		}
	}
}

// Verify if the bearer string is valid
func ValidateBearerString(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	if len(cBearer) == 0 { 
		return false 
	} else { 
		return true 
	}
}

func FormatResponse200( status string, message string, payload string) string {
	p := Payload{} // Create an empty payload struct
	json.Unmarshal([]byte(payload), &p) // 
	response := Response200{
		Status: status,
		Message: message, 
		Payload: p,
	}
	jsonBytes, _ := json.Marshal(response)
	jsonString := string(jsonBytes)
	return jsonString
}

func FormatResponse400( status string, message string) string {
	response := Response400{
		Status: status,
		Message: message, 
	}
	jsonBytes, _ := json.Marshal(response)
	jsonString := string(jsonBytes)
	return jsonString
}
