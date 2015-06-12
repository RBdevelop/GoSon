package apiconv

import "fmt"
import "time"

//gender conversion is easy.... this time
var (
	genderMapping = map[string]string{
		"male":   "M",
		"female": "F",
	}
)

//converting all others 

func Convert(s *SourceData) (TargetData, error) {
	// pre-allocate for len(s.Users)
	t := make(TargetData, 0, len(s.Users))

	for _, sourceUser := range s.Users {
		targetUser := TargetDataUser{
			//sprint method is fast and confusing
			// '' skips what does not need to be converted 
			Address:  fmt.Sprintf("%s\n%s %s %s", sourceUser.Location.Street, sourceUser.Location.City, sourceUser.Location.State, sourceUser.Location.Zip),
			Email:    ``,
			Fullname: ``,
			Gender:   genderMapping[sourceUser.Gender],
			Phone:    sourceUser.Cell,
			Username: ``,
		}
		dob, err := time.Parse("2006-01-02 15:04:05 -0700", sourceUser.DateOfBirth)
		if err != nil {
			return nil, err
		}
		targetUser.DateOfBirth = dob.Format("Monday January 2, 2006")
		t = append(t, targetUser)
	}

	// source data: 1973-06-20 21:28:17 -0500
	// target data: Wednesday June 20, 1973

	// formater template: Mon Jan 2 15:04:05 MST 2006

	return t, nil
}
