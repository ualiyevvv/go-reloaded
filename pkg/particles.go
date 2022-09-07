package goreloaded

import (
	"regexp"
	"strings"
)

func ParticlesFilter(text string) string {
	res := editByA(text)
	res = editByAn(res)
	return res
}

func editByA(text string) string {
	re, err := regexp.Compile(`w+ a\s+\w+ *`)
	checkError(err)
	match := re.FindString(text)
	// fmt.Println("===", match)
	if match != "" {
		word := strings.Split(match, " ")[1]
		if isFromVowel(word) {
			res := strings.Replace(text, match, "an "+word+" ", -1)
			return editByA(res)
		}
	}
	return text
}

func editByAn(text string) string {
	re, err := regexp.Compile(`w+ an\s+\w+ *`)
	checkError(err)
	match := re.FindString(text)
	// fmt.Println("===", match)
	if match != "" {
		word := strings.Split(match, " ")[1]
		if !isFromVowel(word) {
			res := strings.Replace(text, match, "a "+word+" ", -1)
			return editByAn(res)
		}
	}
	return text
}

func isFromVowel(s string) bool {
	vowel := []rune{'a', 'e', 'i', 'o', 'u', 'y'}
	runes := []rune(s)

	for _, v := range vowel {
		if runes[0] == v {
			return true
		}
	}

	return false
}
