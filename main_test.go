package main

import (
	"strings"
	"testing"
)

func concatWithPlus(s1, s2 string) string {
	return s1 + s2
}

func concatWithBuilder(s1, s2 string) string {
	var b strings.Builder
	b.WriteString(s1)
	b.WriteString(s2)
	return b.String()
}

func BenchmarkWithPlus(b *testing.B) {

	for i := 0; i < b.N; i++ {
		concatWithPlus("hello", "world")
	}
}
func BenchmarkWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatWithBuilder("hello", "world")
	}
}
