package dev02

import (
	"errors"
	"strings"
	"unicode"
)

func CheckValidStr(str string) bool {
	beforeNum := false
	runeStr := []rune(str)
	if unicode.IsNumber(runeStr[0]) {
		return false
	}
	for _, i := range runeStr {
		if unicode.IsNumber(i) {
			if beforeNum == true {
				return false
			} else {
				beforeNum = true
			}
		} else {
			beforeNum = false
		}
	}
	return true
}

func UnpackingStr(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	if CheckValidStr(str) == false {
		return "", errors.New("not valid string")
	}
	retStr := strings.Builder{}
	runeStr := []rune(str)
	letterСount := 1
	run := rune(1)
	for _, r := range runeStr {
		if unicode.IsNumber(r) {
			letterСount = int(r-48) - 1
			for i := 0; i < letterСount; i++ {
				retStr.WriteRune(run)
			}
		} else {
			run = r
			retStr.WriteRune(run)
			letterСount = 1
		}
	}
	return retStr.String(), nil
}
