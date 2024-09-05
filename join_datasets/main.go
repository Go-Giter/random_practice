package main

import (
	"fmt"
	"strings"
)

// Write a function that joins two data sets on a key, returning the joined data.
// Bonus points to pretty print the result.

// Example inputs:
// key:  "name"
// data_set_1:  [["name", "ss"], ["john tree", "555-22-5555"], ["amanda plum", "444-11-4444"]]
// data_set_2: [["birthday", "name"], ["2000-01-01", "john tree"], ["1999-02-02", "george stone"]]

// Example output:
// [["name", "ss", "birthday"], ["john tree", "555-22-5555", "2000-01-01"], ["amanda plum", "444-11-4444", "-"], ["george stone", "-", "1999-02-02"]]

// Pretty printed output:
// name         | ss          | birthday
// john tree    | 555-22-5555 | 2000-01-01
// amanda plum  | 444-11-4444 | -
// george stone | -           | 1999-02-02

func joinData(key string, dataSet1, dataSet2 [][]string) [][]string {
	result := make([][]string, len(dataSet1[1:]))

	for ri := range result {
		result[ri] = make([]string, 0, len(dataSet1)+len(dataSet2)/2)
	}

	resultsMap := make(map[string][]string)

	var (
		dataSet1Key int
		dataSet2Key int
	)

	for i, v := range dataSet1[0] {
		if v == key {
			dataSet1Key = i
		}
	}

	for i, v := range dataSet2[0] {
		if v == key {
			dataSet2Key = i
		}
	}

	for _, set := range dataSet1[1:] {
		if _, ok := resultsMap[set[dataSet1Key]]; !ok {
			resultsMap[set[dataSet1Key]] = append(resultsMap[set[dataSet1Key]], set[dataSet1Key+1:]...)

			continue
		}

		resultsMap[set[dataSet2Key]] = append(resultsMap[set[dataSet2Key]], set[:dataSet1Key]...)
	}

	for _, set := range dataSet2[1:] {
		if _, ok := resultsMap[set[dataSet2Key]]; !ok {
			resultsMap[set[dataSet2Key]] = append(resultsMap[set[dataSet2Key]], set[:dataSet2Key]...)

			continue
		}

		resultsMap[set[dataSet2Key]] = append(resultsMap[set[dataSet2Key]], set[:dataSet2Key]...)
	}

	for k, v := range resultsMap {
		if len(v) < 2 {
			if len(strings.Split(v[0], "-")[0]) < 4 {
				v = append(v, "-")
				resultsMap[k] = v

				continue
			}

			v = append([]string{"-"}, v...)
			resultsMap[k] = v
		}
	}

	fmt.Println(resultsMap)

	return result
}

func main() {
	dataSet1 := [][]string{{"name", "ss"}, {"john tree", "555-22-5555"}, {"amanda plum", "444-11-4444"}}
	dataSet2 := [][]string{{"birthday", "name"}, {"2000-01-01", "john tree"}, {"1999-02-02", "george stone"}}
	key := "name"
	joinData(key, dataSet1, dataSet2)

}
