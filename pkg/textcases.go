package goreloaded

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CasesFilter(text string) string {
	res := text

	res = editByUp(res)
	res = editByCap(res)
	res = editByLow(res)
	// res = editByUpSeveral(res)s
	if findTextCaseCommands(res) {
		res = CasesFilter(res)
	}

	return res
}

func findTextCaseCommands(text string) bool {
	re, err := regexp.Compile(`\w+\s+\(\s*(up|cap|low)\s*\)\s*`)
	checkError(err)
	match := re.FindString(text)
	// fmt.Println("RECCCCC", match)
	return match != ""
}

func editByUpSeveral(text string) string {
	re, err := regexp.Compile(`\w+\s+\(\s*up\s*,\s*\d\s*\)\s*`)
	checkError(err)
	match := re.FindString(text)
	loc := re.FindStringIndex(text)
	if match != "" {
		all := strings.Split(match, " ")
		word := all[0]
		numString := parseNumeric(all[2])
		num, err := strconv.Atoi(numString)
		checkError(err)
		fmt.Println("TTTTTT", all[1], num, loc, string(text[708]))

		res := strings.Replace(text, match, strings.ToUpper(word)+" ", -1)
		return editByUp(res)
	}
	return text
}

// func forwardByLoc(text string, loc []int, param int) {
// 	locF := loc
// 	for i := loc[0]; i < loc[1]; i++ {
// 		if text[i] == ' ' {
// 			param--
// 			break
// 		}
// 		text[i] = byte(unicode.ToUpper(rune(text[i])))
// 	}

// }

func editByUp(text string) string {
	re, err := regexp.Compile(`\w+\s+\(\s*up\s*\)\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		word := strings.Split(match, " ")[0]
		res := strings.Replace(text, match, strings.ToUpper(word)+" ", -1)
		return editByUp(res)
	}
	return text
}

func editByLow(text string) string {
	re, err := regexp.Compile(`(\w+)\s*\(\s*low\s*\)\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		word := strings.Split(match, " ")[0]
		res := strings.Replace(text, match, strings.ToLower(word)+" ", -1)
		return editByLow(res)
	}
	return text
}

func editByCap(text string) string {
	re, err := regexp.Compile(`(\w+)\s*\(\s*cap\s*\)\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		word := strings.Split(match, " ")[0]
		res := strings.Replace(text, match, Cap(word)+" ", -1)
		return editByCap(res)
	}
	return text
}

func Cap(word string) string {
	runes := []rune(word)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		return string(runes[0]-32) + word[1:]
	}
	return word
}

func parseNumeric(s string) string {
	res := []rune{}
	for _, r := range s {
		if r >= '0' && r <= '9' {
			res = append(res, r)
		}
	}

	return string(res)

}
