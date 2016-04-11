package main

import (
	"fmt"
	"strings"
	"testing"
)

/*
Chriss-Air:chapter1 quii$ go test -bench .
testing: warning: no tests to run
PASS
BenchmarkJoin-4       	  100000	     21479 ns/op
BenchmarkConcatenate-4	    3000	    509015 ns/op
ok  	github.com/quii/go-book/chapter1	3.972s
*/

func BenchmarkJoin(b *testing.B) {
	fmt.Println("Exercise 1.3 (a)")
	joinConcatenate := func(args []string) string {
		return strings.Join(args, " ")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		joinConcatenate(bigArgs)
	}
}

func BenchmarkConcatenate(b *testing.B) {
	fmt.Println("Exercise 1.3 (b)")

	loopConcatenate := func(args []string) string {
		s, sep := "", ""
		for _, arg := range args {
			s += sep + arg
			sep = " "
		}
		return s
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loopConcatenate(bigArgs)
	}
}
