package util

import (
	"fmt"
	"regexp"
	"strings"
)

const RegPhoneNumber = `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`
const RegEmail = `[a-zA-Z.\-_][a-zA-Z.\-_0-9]{4,}@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z0-9]+\.)+[a-zA-Z]{2,}))`

func IsPhone(phoneNumber string) bool {
	re := regexp.MustCompile(RegPhoneNumber)
	return re.MatchString(phoneNumber)
}

func IsEmail(email string) bool {
	re := regexp.MustCompile(RegEmail)
	return re.MatchString(email)
}

func IsUsername(username string) error {
	if len(username) < 4 {
		return fmt.Errorf("username too short")
	}

	if len(username) > 32 {
		return fmt.Errorf("username too long")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+(?:[_.-][a-zA-Z0-9]+)*$`)
	if !re.MatchString(username) {
		return fmt.Errorf("username wrong format")
	}

	if strings.Contains(username, " ") {
		return fmt.Errorf("username cannot have spaces")
	}

	return nil
}

func IsPassword(password string) error {
	if password == "" {
		return fmt.Errorf("password empty")
	}

	if len(password) < 6 {
		return fmt.Errorf("password too short")
	}

	re := regexp.MustCompile(`^*[0-9]`)
	if !re.MatchString(password) {
		return fmt.Errorf("password must contain number")
	}

	re = regexp.MustCompile(`^*[a-zA-Z]`)
	if !re.MatchString(password) {
		return fmt.Errorf("password must contain letter")
	}

	re = regexp.MustCompile(`^[a-zA-Z0-9 !"#$%&'()*+,-./:;<=>?@[\]^_{|}]*$`)
	if !re.MatchString(password) {
		return fmt.Errorf("invalid character found")
	}

	return nil
}

func IsDisplayName(displayName string) error {
	if strings.HasPrefix(displayName, " ") || strings.HasSuffix(displayName, " ") {
		return fmt.Errorf("cannot have spaces at beginning or ending")
	}

	if displayName == "" {
		return fmt.Errorf("displayname empty")
	}

	if len(displayName) < 4 {
		return fmt.Errorf("displayname too short")
	}

	if len(displayName) > 16 {
		return fmt.Errorf("displayname too long")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+(?:[_.-][a-zA-Z0-9]+)*$`)
	if !re.MatchString(displayName) {
		return fmt.Errorf("displayname wrong format")
	}

	return nil
}