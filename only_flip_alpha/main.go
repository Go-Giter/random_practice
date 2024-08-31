package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := `Dar_@$%ren`
	fmt.Println(doFlip(input))
}

func doFlip(s string) string {
	tmp := ""
	toRet := ""
	var flipped bool

	for _, c := range s {
		if unicode.IsLetter(c) {
			if flipped {
				flipped = false
			}

			tmp += string(c)

			continue
		}

		if !flipped {
			toRet += revStr(tmp)
			tmp = ""
		}

		toRet += string(c)
	}

	toRet += revStr(tmp)

	return toRet
}

func revStr(s string) string {
	rev := ""

	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}

	return rev
}
