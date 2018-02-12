package stringbuilder

import (
	"bytes"
)

// builds one string from an array of strings by making a rune array
func BuildStringWithRunes(strs []string) string {
	length := 0
	for _, str := range strs {
		length += len(str)
	}
	returnString := make([]rune, length)
	pos := 0
	for _, str := range strs {
		for _, char := range str {
			returnString[pos] = char
			pos++
		}
	}
	return string(returnString)
}

// builds one string from an array of strings by using a bytes.Buffer type with runes
func BuildStringWithBytesWriteRune(strs []string) string {
	var buffer bytes.Buffer
	for _, str := range strs {
		for _, char := range str {
			buffer.WriteRune(char)
		}
	}
	return buffer.String()
}

// builds one string from an array of strings by using a bytes.Buffer type with strings
func BuildStringWithBytesWriteString(strs []string) string {
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}

// builds one string from an array of strings by using copy
func BuildStringWithCopy(strs []string) string {
	length := 0
	for _, str := range strs {
		length += len(str)
	}
	buildStr := make([]byte, length)
	pos := 0
	for _, str := range strs {
		pos += copy(buildStr[pos:], str)
	}
	return string(buildStr)
}
