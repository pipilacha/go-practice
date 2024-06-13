package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var reg = regexp.MustCompile(`(?:^\+?1?)[^\d]*([2-9]\d\d)[^\d]*([2-9]\d\d)[^\d]*(\d{4})(?:$|[^\d])`)

func Number(phoneNumber string) (string, error) {
	groups := reg.FindStringSubmatch(phoneNumber)

	if len(groups) != 4 {
		return "", errors.New("invalid phone number")
	}

	return fmt.Sprintf("%s%s%s", groups[1], groups[2], groups[3]), nil
}

func AreaCode(phoneNumber string) (string, error) {
	groups := reg.FindStringSubmatch(phoneNumber)

	if len(groups) != 4 {
		return "", errors.New("invalid area code")
	}

	return groups[1], nil
}

func Format(phoneNumber string) (string, error) {
	groups := reg.FindStringSubmatch(phoneNumber)

	if len(groups) != 4 {
		return "", errors.New("")
	}

	return fmt.Sprintf("(%s) %s-%s", groups[1], groups[2], groups[3]), nil
}
