package reversechar

import "strings"

func ReverseString(str string) string {
	var builder strings.Builder
	builder.Grow(len(str))
	runes := []rune(str)
	
	for i := len(runes) - 1; i >= 0; i-- {
		_, _ = builder.WriteRune(runes[i])	
	}

	return builder.String()
}
