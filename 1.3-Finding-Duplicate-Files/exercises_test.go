package main

import (
	"io"
	"strings"
	"testing"
)

func TestItCountsDuplicatesInFiles(t *testing.T) {
	file1 := `
cats
dogs
cats
  `
	file2 := `
cats
penguins
cats
  `

	files := make(map[string]io.Reader)
	files["file1"] = strings.NewReader(file1)
	files["file2"] = strings.NewReader(file2)

	result := countDuplicates(files)

	if result["cats"].count != 4 {
		t.Error("There should be 4 cats counted across files but got", result["cats"])
	}

	if len(result["cats"].fileNames) != 2 {
		t.Error("Cats should appear in 2 files")
	}
}
