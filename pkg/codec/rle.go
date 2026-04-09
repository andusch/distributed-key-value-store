package codec

import (
	"strconv"
	"strings"
)

type RLECodec struct{}

func (r RLECodec) Encode(s string) string {
	if len(s) == 0 {
		return "0#"
	}

	var result strings.Builder
	cnt := 1
	current := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == current {
			cnt++
		} else {
			result.WriteString(strconv.Itoa(cnt))
			result.WriteByte(current)
			current = s[i]
			cnt = 1
		}
	}

	result.WriteString(strconv.Itoa(cnt))
	result.WriteByte(current)
	return result.String()

}

func (r RLECodec) Decode(s string) string {
	if s == "0#" {
		return ""
	}

	var result strings.Builder
	i := 0

	for i < len(s) {
		count := 0
		hasDigits := false

		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
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

		char := s[i]
		for j := 0; j < count; j++ {
			result.WriteByte(char)
		}
		i++
	}

	return result.String()
}
