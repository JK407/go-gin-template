package common

import (
	"hash/crc32"
	"strings"
)

func CRC32(str string) uint32 {
	str = strings.Replace(str, `/`, `\/`, -1)
	return crc32.ChecksumIEEE([]byte(str))
}

func InArrayForInt(need int, needArr []int) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func InArrayForString(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}
