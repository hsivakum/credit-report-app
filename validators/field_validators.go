package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

func ValidDOB(fl validator.FieldLevel) bool {
	dob := fl.Field().String()

	splitDob := strings.Split(dob, "-")
	if len(splitDob) == 3 {
		if len(splitDob[0]) == 4 {
			if _, err := strconv.Atoi(splitDob[0]); err == nil {
				if len(splitDob[1]) == 2 {
					if _, err := strconv.Atoi(splitDob[1]); err == nil {
						if len(splitDob[2]) == 2 {
							if _, err := strconv.Atoi(splitDob[2]); err == nil {
								return true
							} else {
								fmt.Printf("%q Date should be in right format of DD.\n", splitDob[2])
								return false
							}
						}
					} else {
						fmt.Printf("%q Month should be in right format of MM.\n", splitDob[1])
						return false
					}
				}
			} else {
				fmt.Printf("%q Year should be in right format of YYYY.\n", splitDob[0])
				return false
			}
		}
	}
	return false
}

func ValidSSN(fl validator.FieldLevel) bool {
	dob := fl.Field().String()

	if len(dob) == 9 {
		if _, err := strconv.Atoi(dob); err == nil {
			return true
		}
		return false
	}
	return false
}
