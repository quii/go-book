package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	files := make(map[string]io.Reader)

	fileArgs := os.Args[1:]

	if len(fileArgs) == 0 {
		files["stdin"] = os.Stdin
	} else {
		for _, arg := range fileArgs {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			files[arg] = f
			defer f.Close()
		}
	}

	for line, report := range countDuplicates(files) {
		if report.count > 1 {
			fmt.Printf("%s - %d - %v\n", line, report.count, report.fileNames)
		}
	}
}

type report struct {
	count     int
	fileNames []string
}

func countDuplicates(files map[string]io.Reader) map[string]report {
	lineReports := make(map[string]report)

	for fileName, rdr := range files {

		for line, count := range countLines(rdr) {

			lineReport, exists := lineReports[line]
			if !exists {
				lineReport = report{
					count:     count,
					fileNames: []string{fileName},
				}
			} else {
				lineReport.count += count
				lineReport.fileNames = append(lineReport.fileNames, fileName)
			}

			lineReports[line] = lineReport
		}
	}
	return lineReports
}

func countLines(f io.Reader) map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts
}
