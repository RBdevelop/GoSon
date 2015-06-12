package apiconv

// first thing first convert Json to Golang struct aka GoSon  
// reference https://gobyexample.com/json

import (
	"encoding/json"
	"os"
)

// SourceData is structure for the ServiceA data
// break the json into intelligible sections this time 
type SourceData struct {
	Count uint             `json:"count"`
	Users []SourceDataUser `json:"users"`
}

// SourceDataUser holds  ServiceA info  per  single user in struct format
type SourceDataUser struct {
	Gender string `json:"gender"`
	Name   struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Location struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
		Zip    string `json:"zip"`
	} `json:"location"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Registered  string `json:"registered"` //++ needs custom  parsing . will ask John 
	DateOfBirth string `json:"dob"`        //++ needs custom  parsing 
	Phone       string `json:"phone"`
	Cell        string `json:"cell"`
	Picture     struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	} `json:"picture"`
	SocialSecurityNumber string `json:"ssn"`
}

func ReadSourceFile(filename string) (*SourceData, error) {
	// open ze file for reading
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// decode file contents because the files are in the computer!  
	data := &SourceData{}
	err = json.NewDecoder(f).Decode(data)
	if err != nil {
		return nil, err
	}

	// should return data or err but to err is human 
	return data, nil
}

