package core

import (
	"fmt"
	"regexp"
)

func ValidateChinesePhoneNumber(phoneNumber string) bool {
	// 使用正则表达式匹配中国手机号的格式
	regex := `^1[3456789]\d{9}$`
	match, err := regexp.MatchString(regex, phoneNumber)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return match
}
