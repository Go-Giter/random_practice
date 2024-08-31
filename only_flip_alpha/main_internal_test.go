package main

import (
	"testing"

	"github.com/reedobrien/checkers"
)

func TestDoFlip(t *testing.T) {
	input := "Th)(*^)isATest"
	got := doFlip(input)

	checkers.Equals(t, got, "hT)(*^)tseTAsi")
}
