package utils

import "time"

func IsAValidBirthDate(date time.Time) bool {
	now := time.Now()

	years := now.Year() - date.Year()

	return years < 80 && years > 14
}
