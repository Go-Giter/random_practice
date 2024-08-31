package main

import (
	"testing"

	"github.com/reedobrien/checkers"
)

func TestTracker(t *testing.T) {

	table := []struct {
		name       string
		testString string
		want       bool
	}{
		{
			"should return true ()",
			`()`,
			true,
		},
		{
			"should return true {[()]}",
			`{[()]}`,
			true,
		},
		{
			"should return false (]",
			`(]`,
			false,
		},
		{
			"should return false {])",
			`{])`,
			false,
		},
		{
			"should return false {()",
			`{()`,
			false,
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			got := tracker(tc.testString)
			checkers.Equals(t, got, tc.want)
		})
	}

}
