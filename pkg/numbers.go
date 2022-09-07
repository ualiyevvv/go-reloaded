package goreloaded

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func NumFilter(text string) string {
	res := editByHex(text)
	res = editByBin(res)
	res = editByEmpty(res)
	return res
}

func editByEmpty(text string) string {
	reHex, err := regexp.Compile(`\s+\(\s*(hex|bin)\s*\)\s*`)
	checkError(err)
	matchHex := reHex.FindString(text)
	if matchHex != "" {
		res := strings.Replace(text, matchHex, " ", -1)
		return editByEmpty(res)
	}

	return text
}

func editByHex(text string) string {
	reHex, err := regexp.Compile(`\s+(\d+)\s*\(\s*hex\s*\)\s*`)
	checkError(err)
	match := reHex.FindString(text)
	if match != "" {
		num := strings.Split(match, " ")[1]
		fmt.Println("@@@@", num)
		res := strings.Replace(text, match, " "+ChangeNumRegister(num, 16)+" ", -1)
		return editByHex(res)
	}
	return text
}

func editByBin(text string) string {
	reBin, err := regexp.Compile(`\s+(\d+)\s*\(\s*bin\s*\)\s*`)
	checkError(err)
	match := reBin.FindString(text)
	if match != "" {
		num := strings.Split(match, " ")[1]
		res := strings.Replace(text, match, " "+ChangeNumRegister(num, 2)+" ", -1)
		return editByBin(res)
	}
	return text
}

func ChangeNumRegister(word string, base int) string {
	res, err := strconv.ParseInt(word, base, 32)
	checkError(err)

	return strconv.Itoa(int(res))
}
