package goreloaded

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	signs = []string{"!", ":", ";", ",", "[...]"}
	// multiSigns = []string{"...", "?!"}
	// quotes     = []rune{'\'', '"', '`', '\''}
)

func PunctuationsFilter(text string) string {
	var res string
	// for _, r := range signs {
	res = editByComma(text)
	res = editByDComma(res)
	res = editByV(res)
	res = editByV2(res)
	res = editByV3(res)
	res = editByV4(res)
	res = editBySpace(res)
	// res = editByQuote2(res)
	res = editQuote(res)

	// }

	return res
}

func editByComma(text string) string {
	re, err := regexp.Compile(`\s+,\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, ", ", -1)
		return editByComma(res)
	}
	return text
}

var count = 0

func editByQuote(text string) string {
	r := editQuote(text)

	// arrtxt := strings.Split(text, " ")
	// newtxt := ""
	// fmt.Println(text)
	// for _, r := range arrtxt {
	// 	if r == "" {
	// 		return ""
	// 	}
	// 	r = editQuote(r)
	// 	newtxt += r
	// }
	return r
}

func editQuote(s string) string {
	countQ := 0
	fcase := false
	newtxt := ""
	fmt.Println(s)
	for i, r := range s {
		if fcase {
			fcase = false
			continue
		}
		if r == '"' && countQ%2 == 0 {
			if i != 0 && s[i-1] != ' ' {
				newtxt += ` "`
				if i != len(s)-1 && s[i+1] == ' ' {
					fcase = true
				}
				if i != len(s)-1 && s[i+1] != ' ' {
					newtxt += ` "`
				}
				countQ++
				continue
			}
		}
		// if fcase {
		// 	fcase = false
		// 	continue
		// }
		// if r == ' ' && i != 0 && s[i-1] == '"' && countQ%2 == 0 {
		// 	countQ++
		// 	continue
		// }
		// if r == '"' && s[i+1] == ' ' && s[i-1] != ' ' && i != 0 && countQ%2 == 0 && !fcase {
		// 	newtxt += ` "`
		// 	countQ++
		// 	fcase = true
		// 	continue
		// }
		// if r == ' ' && i != len(s)-1 && s[i+1] == '"' && countQ%2 != 0 {
		// 	countQ++
		// 	continue
		// }
		newtxt += string(r)
	}
	fmt.Println(newtxt)
	return newtxt
}

func editByQuote2(text string) string {
	re, err := regexp.Compile(`\s+'`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, "'", -1)
		return editByQuote2(res)
	}
	return text
}

func editBySpace(text string) string {
	re, err := regexp.Compile(` \s+`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, " ", -1)
		return editBySpace(res)
	}
	return text
}

func editByDComma(text string) string {
	re, err := regexp.Compile(`\s+:\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, ": ", -1)
		return editByDComma(res)
	}
	return text
}
func editByV(text string) string {
	re, err := regexp.Compile(`\s+\?\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, "? ", -1)
		return editByV(res)
	}
	return text
}
func editByV2(text string) string {
	re, err := regexp.Compile(`\s+!\?\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, "!? ", -1)
		return editByV2(res)
	}
	return text
}
func editByV3(text string) string {
	re, err := regexp.Compile(`\s+\.\.\.\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, "... ", -1)
		return editByV3(res)
	}
	return text
}
func editByV4(text string) string {
	re, err := regexp.Compile(`\s+\.\s*`)
	checkError(err)
	match := re.FindString(text)
	if match != "" {
		res := strings.Replace(text, match, ". ", -1)
		return editByV4(res)
	}
	return text
}

// func editByEllipsis(text string) string {
// 	re, err := regexp.Compile(`\s+...\s+`)
// 	checkError(err)
// 	match := re.FindString(text)
// 	if match != "" {
// 		res := strings.Replace(text, match, ". ", -1)
// 		return editByEllipsis(res)
// 	}
// 	return text
// }
