package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var buf strings.Builder
	for _, n := range phoneNumber {
		if n < '0' || n > '9' {
			continue
		}
		buf.WriteRune(n)
	}
	if buf.Len() > 11 || buf.Len() < 10 {
		return "", errors.New("wrong phone number format")
	}
	if buf.Len() == 11 {
		result := buf.String()
		if result[0] == '1' && result[1] > '1' && result[4] > '1' {
			return result[1:], nil
		}
		return "", errors.New("wrong phohe number format")
	}
	result := buf.String()
	if result[0] < '2' || result[3] < '2' {
		return "", errors.New("wrong phone number format")
	}
	return result, nil
}

func AreaCode(phoneNumber string) (string, error) {
	result, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return result[:3], nil
}

func Format(phoneNumber string) (string, error) {
	result, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", result[:3], result[3:6], result[6:10]), nil
}
