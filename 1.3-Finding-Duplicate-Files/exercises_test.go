package main

import (
	"strings"
	"testing"
)

func TestItCountsLinesOfTest(t *testing.T) {
	text := `
cats
dogs
cats
  `
	input := strings.NewReader(text)

	results := make(map[string]int)

	countLines(input, results)

	if results["cats"] != 2 {
		t.Error("Should be 2 cats")
	}

	if results["dogs"] != 1 {
		t.Error("Should be 1 dog")
	}
}
