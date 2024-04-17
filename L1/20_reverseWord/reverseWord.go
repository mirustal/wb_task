package reverseword

import "strings"

func ReverseWords(str string) string {
	words := strings.Fields(str)
	var builder strings.Builder

	builder.Grow(len(str))

	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		if i > 0 {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}