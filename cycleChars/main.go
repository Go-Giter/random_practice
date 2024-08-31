package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"unicode"
)

func cycle(s string, indexToFlip int, listCh chan<- string) {
	if indexToFlip > len(s)-1 || indexToFlip < 0 {
		listCh <- strings.ToLower(s)
		listCh <- strings.ToUpper(s)
		close(listCh)

		return
	}

	ss := strings.Split(s, "")

	nc := flipChar(ss[indexToFlip])

	ss[indexToFlip] = nc

	listCh <- strings.Join(ss, "")

	cycle(s, indexToFlip+1, listCh)
}

func flipChar(s string) string {
	if unicode.IsLower([]rune(s)[0]) {
		return strings.ToUpper(s)
	}

	return strings.ToLower(s)
}

func main() {
	input := "DarrenTerry"
	listCh := make(chan string, len(input)+2)
	cycle(strings.ToLower(input), 0, listCh)
	result := make([]string, 0, len(listCh))

	for s := range listCh {
		result = append(result, s)
	}

	slices.SortStableFunc[[]string](result, func(a, b string) int {
		return cmp.Compare[string](a, b)
	})

	fmt.Printf("%#v\n", result)
}
