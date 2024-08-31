package main

import "fmt"

func main() {
	fmt.Println(tracker(`{()`))
}

func tracker(input string) bool {
	tracker := []rune{}

	for _, c := range input {
		switch c {
		case '(':
			tracker = append(tracker, c)
		case '[':
			tracker = append(tracker, c)
		case '{':
			tracker = append(tracker, c)

		default:
			switch c {
			case ')':
				if tracker[len(tracker)-1] != '(' {
					return false
				}
			case ']':
				if tracker[len(tracker)-1] != '[' {
					return false
				}
			case '}':
				if tracker[len(tracker)-1] != '{' {
					return false
				}
			}

			tracker = tracker[:len(tracker)-1]
		}
	}

	return len(tracker) == 0
}
