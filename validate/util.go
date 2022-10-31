package util

import "regexp"

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
