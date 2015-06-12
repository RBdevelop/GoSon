package apiconv

import (
	"encoding/json"
	"os"
)

//this is our target / bullseye 
// {
//   "address": "3258 Thornridge Cir\nFrederick, Texas  17354",
//   "date_of_birth": "Wednesday June 20, 1973",
//   "email": "henry.howard66@example.com",
//   "fullname": "Mr. Henry Howard",
//   "gender": "M",
//   "phone": "(880)-850-2345",
//   "username": "tinyleopard307"
// }

// TargetData is an array of user information for serviceB
type TargetData []TargetDataUser

// TargetDataUser contains the user information structured for serviceB
type TargetDataUser struct {
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	Username    string `json:"username"`
}

func WriteTargetData(t TargetData) error {
	err := json.NewEncoder(os.Stdout).Encode(t)
	if err != nil {
		return err
	}
	return nil
}
