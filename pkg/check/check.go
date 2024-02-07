package check

import (
	"errors"
	"fmt"
	"time"
)

func CalculateAge(birthDate string) int {
	layout := "2006-01-02"
	birthday, err := time.Parse(layout, birthDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return 0
	}

	now := time.Now()
	age := now.Year() - birthday.Year()

	if now.YearDay() < birthday.YearDay() {
		age--
	}

	return age
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password length should be more than 6")
	}

	return nil
}
