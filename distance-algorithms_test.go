package main

import "testing"

func runTest(b *testing.B, name string, fn func(a, b string) int, expectedDistance int, argA, argB string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := fn(argA, argB); got != expectedDistance {
				b.Error("bad distance")
			}
		}
	})
}

func BenchmarkExtraShort(b *testing.B) {
	runTest(b, "Exponential", LevenshteinDistanceExponential, 0, "", "")
	runTest(b, "FullMatrix", LevenshteinDistanceFullMatrix, 0, "", "")
	runTest(b, "TwoRow", LevenshteinDistanceTwoRow, 0, "", "")
	runTest(b, "OneRow", LevenshteinDistanceOneRow, 0, "", "")
}

func BenchmarkVeryShort(b *testing.B) {
	runTest(b, "Exponential", LevenshteinDistanceExponential, 0, "a", "a")
	runTest(b, "FullMatrix", LevenshteinDistanceFullMatrix, 0, "a", "a")
	runTest(b, "TwoRow", LevenshteinDistanceTwoRow, 0, "a", "a")
	runTest(b, "OneRow", LevenshteinDistanceOneRow, 0, "a", "a")
}

func BenchmarkShort(b *testing.B) {
	runTest(b, "Exponential", LevenshteinDistanceExponential, 0, "abc", "abc")
	runTest(b, "FullMatrix", LevenshteinDistanceFullMatrix, 0, "abc", "abc")
	runTest(b, "TwoRow", LevenshteinDistanceTwoRow, 0, "abc", "abc")
	runTest(b, "OneRow", LevenshteinDistanceOneRow, 0, "abc", "abc")
}

func BenchmarkMedium(b *testing.B) {
	runTest(b, "Exponential", LevenshteinDistanceExponential, 0, "abcdef", "abcdef")
	runTest(b, "FullMatrix", LevenshteinDistanceFullMatrix, 0, "abcdef", "abcdef")
	runTest(b, "TwoRow", LevenshteinDistanceTwoRow, 0, "abcdef", "abcdef")
	runTest(b, "OneRow", LevenshteinDistanceOneRow, 0, "abcdef", "abcdef")
}

func BenchmarkLong(b *testing.B) {
	runTest(b, "Exponential", LevenshteinDistanceExponential, 0, "abcdefabcdef", "abcdefabcdef")
	runTest(b, "FullMatrix", LevenshteinDistanceFullMatrix, 0, "abcdefabcdef", "abcdefabcdef")
	runTest(b, "TwoRow", LevenshteinDistanceTwoRow, 0, "abcdefabcdef", "abcdefabcdef")
	runTest(b, "OneRow", LevenshteinDistanceOneRow, 0, "abcdefabcdef", "abcdefabcdef")
}

func BenchmarkVeryLong(b *testing.B) {
	// runTest(b, "Exponential", LevenshteinDistanceExponential, 0, "abcdefabcdefabcdefabcdef", "abcdefabcdefabcdefabcdef") // block everything
	runTest(b, "FullMatrix", LevenshteinDistanceFullMatrix, 0, "abcdefabcdefabcdefabcdef", "abcdefabcdefabcdefabcdef")
	runTest(b, "TwoRow", LevenshteinDistanceTwoRow, 0, "abcdefabcdefabcdefabcdef", "abcdefabcdefabcdefabcdef")
	runTest(b, "OneRow", LevenshteinDistanceOneRow, 0, "abcdefabcdefabcdefabcdef", "abcdefabcdefabcdefabcdef")
}
