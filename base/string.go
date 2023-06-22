package base

import (
	"regexp"
	"strings"
)

func RemoveVietnameseAccents(pattern string) string {

	var codau = []string{"À", "Á", "Â", "Ã", "È", "É",
		"Ê", "Ì", "Í", "Ò", "Ó", "Ô", "Õ", "Ù", "Ú", "Ý", "Ỳ", "Ỷ", "Ỵ", "Ỹ",
		"à", "á", "â", "ã", "è", "é", "ê", "ì", "í", "ò", "ó", "ô", "õ",
		"ù", "ú", "ý", "ỳ", "ỷ", "ỵ", "ỹ", "Ă", "ă", "Đ", "đ", "Ĩ", "ĩ", "Ũ",
		"ũ", "Ơ", "ơ", "Ư", "ư", "Ạ", "ạ", "Ả", "ả", "Ấ", "ấ", "Ầ", "ầ",
		"Ẩ", "ẩ", "Ẫ", "ẫ", "Ậ", "ậ", "Ắ", "ắ", "Ằ", "ằ", "Ẳ", "ẳ", "Ẵ",
		"ẵ", "Ặ", "ặ", "Ẹ", "ẹ", "Ẻ", "ẻ", "Ẽ", "ẽ", "Ế", "ế", "Ề", "ề",
		"Ể", "ể", "Ễ", "ễ", "Ệ", "ệ", "Ỉ", "ỉ", "Ị", "ị", "Ọ", "ọ", "Ỏ",
		"ỏ", "Ố", "ố", "Ồ", "ồ", "Ổ", "ổ", "Ỗ", "ỗ", "Ộ", "ộ", "Ớ", "ớ",
		"Ờ", "ờ", "Ở", "ở", "Ỡ", "ỡ", "Ợ", "ợ", "Ụ", "ụ", "Ủ", "ủ", "Ứ",
		"ứ", "Ừ", "ừ", "Ử", "ử", "Ữ", "ữ", "Ự", "ự"}

	// Mang cac ky tu thay the khong dau
	var khongdau = []string{"A", "A", "A", "A", "E",
		"E", "E", "I", "I", "O", "O", "O", "O", "U", "U", "Y", "Y", "Y", "Y",
		"Y", "a", "a", "a", "a", "e", "e", "e", "i", "i", "o", "o", "o",
		"o", "u", "u", "y", "y", "y", "y", "y", "A", "a", "D", "d", "I", "i",
		"U", "u", "O", "o", "U", "u", "A", "a", "A", "a", "A", "a", "A",
		"a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a",
		"A", "a", "A", "a", "E", "e", "E", "e", "E", "e", "E", "e", "E",
		"e", "E", "e", "E", "e", "E", "e", "I", "i", "I", "i", "O", "o",
		"O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O",
		"o", "O", "o", "O", "o", "O", "o", "O", "o", "U", "u", "U", "u",
		"U", "u", "U", "u", "U", "u", "U", "u", "U", "u"}

	for i, charcodau := range codau {
		pattern = strings.ReplaceAll(pattern, charcodau, khongdau[i])
	}
	return pattern
}

func RemoveSpecialCharacters(pattern string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return ""
	}
	processedString := reg.ReplaceAllString(RemoveVietnameseAccents(pattern), "")
	return processedString
}
