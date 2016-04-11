package main

import (
	"bytes"
	"strings"
	"testing"
)

var bigArgs []string

func init() {
	for i := 0; i < 10000; i++ {
		bigArgs = append(bigArgs, "x")
	}
}

/*
Chriss-Air:chapter1 quii$ go test -bench .
testing: warning: no tests to run
PASS
BenchmarkJoin-4       	  100000	     21479 ns/op
BenchmarkConcatenate-4	    3000	    509015 ns/op
ok  	github.com/quii/go-book/chapter1	3.972s
*/

// Exercise 1.3
func BenchmarkJoin(b *testing.B) {
	joinConcatenate := func(args []string) string {
		return strings.Join(args, " ")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		joinConcatenate(bigArgs)
	}
}

// Exercise 1.3
func BenchmarkConcatenate(b *testing.B) {
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

// Exercise 1.3 extra credit?
//
func BenchmarkBuffer(b *testing.B) {
	bufferConcat := func(args []string) string {
		var buffer bytes.Buffer
		sep := ""
		for _, arg := range args {
			buffer.WriteString(sep)
			buffer.WriteString(arg)
			sep = " "
		}
		return buffer.String()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bufferConcat(bigArgs)
	}
}
