package main

import "fmt"

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
	result := make([][]string, len(dataSet1)+len(dataSet2)/2)

	for ri := range result {
		result[ri] = make([]string, 0, len(dataSet1)+len(dataSet2)/2)
	}

	joinPoint := make(map[string]int)
	joinedData := make(map[string][]string)

	for i, v := range dataSet1[0] {
		if v == key {
			joinPoint["dataSet1"] = i
		}
	}

	for i, v := range dataSet2[0] {
		if v == key {
			joinPoint["dataSet2"] = i
		}
	}

	if len(joinPoint) > 1 {
		for _, set := range dataSet1[1:] {
			k := set[joinPoint["dataSet1"]]
			if _, ok := joinedData[k]; !ok {
				joinedData[k] = make([]string, 0)
				set = append(set[:joinPoint["dataSet1"]], set[joinPoint["dataSet1"]+1:]...)
				joinedData[k] = append(joinedData[k], set...)

				continue
			}
			set = append(set[:joinPoint["dataSet2"]], set[joinPoint["dataSet2"]+1:]...)
			joinedData[k] = append(joinedData[k], set...)
		}

		for _, set := range dataSet2[1:] {
			k := set[joinPoint["dataSet2"]]
			if _, ok := joinedData[k]; !ok {
				joinedData[k] = make([]string, 0)
				set = append(set[:joinPoint["dataSet2"]], set[joinPoint["dataSet2"]+1:]...)
				joinedData[k] = append(joinedData[k], set...)

				continue
			}
			set = append(set[:joinPoint["dataSet2"]], set[joinPoint["dataSet2"]+1:]...)
			joinedData[k] = append(joinedData[k], set...)
		}
	}

	result[0] = []string{"name", "ss", "birthday"}

	count := 1
	for k, v := range joinedData {
		result[count] = append(result[count], k)
		result[count] = append(result[count], v...)
		count++
	}

	return result
}

func main() {
	dataSet1 := [][]string{{"name", "ss"}, {"john tree", "555-22-5555"}, {"amanda plum", "444-11-4444"}}
	dataSet2 := [][]string{{"birthday", "name"}, {"2000-01-01", "john tree"}, {"1999-02-02", "george stone"}}
	key := "name"
	fmt.Println(joinData(key, dataSet1, dataSet2))
}
