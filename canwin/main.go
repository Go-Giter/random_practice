package main

func canWin(board []int, startingIndex int, tracker map[int]struct{}) bool {
	if startingIndex > len(board)-1 || startingIndex < 0 {
		return false
	}

	if _, ok := tracker[startingIndex]; ok {
		return !ok
	}

	tracker[startingIndex] = struct{}{}

	jumpVal := board[startingIndex]

	if jumpVal == 0 {
		return true
	}

	return canWin(board, startingIndex+jumpVal, tracker) || canWin(board, startingIndex-jumpVal, tracker)
}

func main() {
	runTest([]int{1, 1}, 0, false)
	runTest([]int{0, 4, 2, 1}, 2, true)
	runTest([]int{1, 10, 2, 4}, 3, false)
	runTest([]int{1, 0, 3, 4, 0}, 1, true)
}

func runTest(board []int, startingIndex int, want bool) {
	tracker := make(map[int]struct{})

	if canWin(board, startingIndex, tracker) != want {
		panic("mismatched")
	}
}
