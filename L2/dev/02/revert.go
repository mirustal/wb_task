package revert

import (
	"strconv"
	"unicode"
)

func RLERevert(str string) string {
	data := []rune(str)

	var result []rune
	i := 0
	for i < len(data) {
		if unicode.IsDigit(data[i]) {
			number := make([]rune, 0)
			for i < len(data) && unicode.IsDigit(data[i]) {
				number = append(number, data[i])
				i++
			}
			n, err := strconv.Atoi(string(number))
			if err != nil {
				return ""
			}

			for j := 0; j < n-1; j++ {
				result = append(result, result[len(result)-1])
			}
		} else {
			result = append(result, data[i])
		}
		i++
	}

	return string(result)
}
