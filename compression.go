package main

import (
	"reflect"
	"strconv"
	"strings"
)

func run_length_encode(s string) string {

	// String verification
	aux := reflect.TypeOf(s).Kind()
	if aux != reflect.String {
		return ""
	}

	var result string
	cnt := 1
	current_char := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == current_char {
			cnt += 1
		} else {
			result += strconv.Itoa(cnt)
			result += string(current_char)
			current_char = s[i]
			cnt = 1
		}
	}

	result += strconv.Itoa(cnt)
	result += string(current_char)

	return "" + result

}

func run_length_decode(s string) string {

	var result strings.Builder

	i := 0

	for i < len(s) {

		count := 0
		hasDigits := false

		for i < len(s) && '0' <= s[i] && s[i] <= '9' {
			hasDigits = true
			count = count*10 + int(s[i]-'0')
			i++
		}

		if !hasDigits {
			count = 1
		}

		if i >= len(s) {
			break
		}

		aux := s[i]
		for j := 0; j < count; j++ {
			result.WriteByte(aux)
		}

		i++

	}

	return result.String()

}
