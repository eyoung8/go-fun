package stringalgos

import "testing"

func TestIsUniqueTrueCase(t *testing.T) {
	str := "abcdefg"
	if unique := IsUnique(str); !unique {
		t.Errorf("unique string=%s was not found to be unique!", str)
	}
}

func TestIsUniqueFalseCase(t *testing.T) {
	str := "abcdefgsdfsafs"
	if unique := IsUnique(str); unique {
		t.Errorf("non-unique string=%s was found to be unique!", str)
	}
}

func TestIsPermutationTrueCase(t *testing.T) {
	strA := "abcdefg"
	strB := "gdefcba"
	if permutation := IsPermutation(strA, strB); !permutation {
		t.Errorf("string: %s was incorrectly found not to be permutation of string: %s", strA, strB)
	}
}

func TestIsPermutationFalseCase(t *testing.T) {
	strA := "abcdefg"
	strB := "zzzzada"
	if permutation := IsPermutation(strA, strB); permutation {
		t.Errorf("string: %s was incorrectly found to be permutation of string: %s", strA, strB)
	}
}

func TestIsPermutationDifferentLengths(t *testing.T) {
	strA := "abcdefg"
	strB := "zzzzadaasdad"
	if permutation := IsPermutation(strA, strB); permutation {
		t.Errorf("string: %s was incorrectly found to be permutation of string: %s", strA, strB)
	}
}

func TestURLifyWithSpaces(t *testing.T) {
	str := "This is a sentence."
	expected := "This%20is%20a%20sentence."
	if got := URLify(str); got != expected {
		t.Errorf("Got: %s, Expected: %s", got, expected)
	}
}

func TestURLifyNoSpaces(t *testing.T) {
	str := "Thisisasentence."
	expected := "Thisisasentence."
	if got := URLify(str); got != expected {
		t.Errorf("Got: %s, Expected: %s", got, expected)
	}
}

func TestPalindromePermutation(t *testing.T) {
	cases := []struct {
		str      string
		expected bool
	}{
		{"taco cat", true},
		{"atco cta", true},
		{"abcde fg", false},
	}

	for _, c := range cases {
		if got := PalindromePermutation(c.str); got != c.expected {
			t.Errorf("Case: %s: Got: %v, Expected: %v", c.str, got, c.expected)
		}
	}
}

func TestOneWay(t *testing.T) {
	cases := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"toolong", "too", false},
		{"pale", "ale", true},
		{"ale", "pale", true},
		{"pae", "pale", true},
		{"ple", "pale", true},
		{"pale", "pale", true},
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "pales", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
	}

	for _, c := range cases {
		if got := OneWay(c.str1, c.str2); got != c.expected {
			t.Errorf("Case: %s -> %s: Got: %v, Expected: %v", c.str1, c.str2, got, c.expected)
		}
	}
}

func TestCompress(t *testing.T) {
	cases := []struct {
		str      string
		expected string
	}{
		{"abc", "abc"},
		{"aaabc", "aaabc"},
		{"aaabbc", "aaabbc"},
		{"aaabbbc", "a3b3c1"},
		{"abbbbbcc", "a1b5c2"},
	}

	for _, c := range cases {
		if got := Compress(c.str); got != c.expected {
			t.Errorf("Case: %s -> %s: Got: %v, Expected: %v", c.str, c.expected, got, c.expected)
		}
	}
}

func TestZeroMatrix(t *testing.T) {
	cases := []struct {
		m        [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3, 4, 0},
				{1, 1, 0, 1, 2},
				{2, 3, 4, 5, 6},
			},
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{2, 3, 0, 5, 0},
			},
		},
	}
	for _, c := range cases {
		matrixCopy := copyMatrix(c.m)
		if got := ZeroMatrix(c.m); !matrixEqual(got, c.expected) {
			t.Errorf("Case: %v -> %v: Got: %v, Expected: %v", matrixCopy, c.expected, got, c.expected)
		}
	}
}

func copyMatrix(m [][]int) [][]int {
	m2 := make([][]int, len(m))
	for _, row := range m {
		newRow := make([]int, len(row))
		copy(newRow, row)
		m2 = append(m2, newRow)
	}
	return m2
}

func matrixEqual(m1 [][]int, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i, row := range m1 {
		for j, col := range row {
			if col != m2[i][j] {
				return false
			}
		}
	}
	return true
}
