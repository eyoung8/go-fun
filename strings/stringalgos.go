package stringalgos

import (
	"bytes"
	"fmt"
	"github.com/eyoung8/go-fun/strings/stringbuilder"
	"sort"
)

// a string is unique if each character is only present once
func IsUnique(s string) bool {
	historyMap := make(map[rune]bool)
	for _, char := range s {
		if _, ok := historyMap[char]; ok {
			return false
		}
		historyMap[char] = true
	}
	return true
}

// string a is a permutation of string b if they have the exact same characters regardless of
// order
func IsPermutation(strA string, strB string) bool {
	if len(strA) != len(strB) {
		return false
	}
	strA = sortString(strA)
	strB = sortString(strB)
	return strA == strB
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Sort(sortRunes(runes))
	return string(runes)
}

// Replaces spaces with %20 to be used in a URL
func URLify(str string) string {
	spaceStr := " "
	return replaceString(str, rune(spaceStr[0]), "%20")
}

func replaceString(str string, replace rune, replaceWith string) string {
	buildSlice := make([]string, len(str))
	for _, char := range str {
		if char == replace {
			buildSlice = append(buildSlice, replaceWith)
		} else {
			buildSlice = append(buildSlice, string(char))
		}
	}
	return stringbuilder.BuildStringWithCopy(buildSlice)
}

// Returns true if the input string has a permutation that is a palindrome
// Ignores spaces but other special characters aren't supported.
func PalindromePermutation(str string) bool {
	charCount := make(map[rune]int)
	for _, char := range str {
		charCount[char] += 1
	}
	delete(charCount, ' ') // remove spaces

	oddCount := false // only one character can have an odd count
	for _, count := range charCount {
		if (count % 2) != 0 {
			if oddCount {
				return false
			}
			oddCount = true
		}
	}
	return true
}

// Returns true if zero or one change to a string will make it equal to the other
// Change = change 1 character or add 1 character
func OneWay(str1 string, str2 string) bool {
	str1len := len(str1)
	str2len := len(str2)
	if abs(str2len-str1len) > 1 {
		return false
	}
	for i, _ := range str1 {
		if i < str2len {
			if str1[i] != str2[i] {
				if str1len > str2len {
					return removeCase(str1, str2, i)
				} else if str1len < str2len {
					return addCase(str1, str2, i)
				} else if str1len == str2len {
					return changeCase(str1, str2, i)
				}
			}
		}
	}
	return true // equal case
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// bale --> baale
func addCase(str1 string, str2 string, i int) bool {
	if len(str1)-1 == i {
		return false
	}
	addStr := str1[:i+1] + string(str2[i]) + str1[i+1:]
	return addStr == str2
}

// bale --> ble
func removeCase(str1 string, str2 string, i int) bool {
	if len(str1)-1 == i {
		return false
	}
	removeStr := str1[:i] + str1[i+1:]
	return removeStr == str2
}

// bale --> pale
func changeCase(str1 string, str2 string, i int) bool {
	if len(str1)-1 == i {
		return false
	}
	changeStr := str1[:i] + string(str2[i]) + str1[i+1:]
	return changeStr == str2
}

// basic string compression algorithm removing repeated characters and replacing with an int
// ex: aaaaabbbbcccc --> a5b4c4
// returns the original string if it can't be shortened
// Only accepts alphabetic characters
func Compress(str string) string {
	char := '0'
	count := 0
	var compressed bytes.Buffer
	for _, c := range str {
		if c != char && count > 0 {
			compressed.WriteString(fmt.Sprintf("%s%d", string(char), count))
			count = 0
		}
		char = c
		count += 1
	}
	compressed.WriteString(fmt.Sprintf("%s%d", string(char), count))
	if len(compressed.String()) >= len(str) {
		return str
	}
	return compressed.String()
}

// Receives 2 dimensional matrix and zeroes out any row/column that contains a 0
// Not done yet
func ZeroMatrix(matrix [][]int) [][]int {
	zeroRows := make(map[int]bool)
	zeroColumns := make(map[int]bool)

	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				zeroRows[i] = true
				zeroColumns[j] = true
			}
		}
	}
	for row, _ := range zeroRows {
		zeroRow(row, matrix)
	}
	for col, _ := range zeroColumns {
		zeroColumn(col, matrix)
	}
	return matrix
}

func zeroRow(row int, matrix [][]int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[row][i] = 0
	}
}

func zeroColumn(col int, matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}
