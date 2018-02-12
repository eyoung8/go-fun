package stringbuilder

import (
	"testing"
)

var (
	testStr = []string{"abcd", "bdec", "ab", "a", "npc", "lol", "haha"}
)

func BenchmarkRunes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildStringWithRunes(testStr)
	}
}

func BenchmarkBytesWriteRune(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildStringWithBytesWriteRune(testStr)
	}
}

func BenchmarkBytesWriteString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildStringWithBytesWriteString(testStr)
	}
}

func BenchmarkCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildStringWithCopy(testStr)
	}
}
