package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
)

func init() {
	flag.IntVar(&after, "A", 0, "Печатать N строк после совпадения")
	flag.IntVar(&before, "B", 0, "Печатать N строк до совпадения")
	flag.IntVar(&context, "C", 0, "Печатать N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "Подсчет количества строк с совпадениями")
	flag.BoolVar(&ignoreCase, "i", false, "Игнорировать регистр при сравнении")
	flag.BoolVar(&invert, "v", false, "Выбрать строки, не содержащие совпадений")
	flag.BoolVar(&fixed, "F", false, "Точное совпадение со строкой")
	flag.BoolVar(&lineNum, "n", false, "Печать номера строки с совпадением")
	flag.Parse()

	if context != 0 {
		after = context
		before = context
	}
	pattern = flag.Arg(0)
}

func main() {
	if pattern == "" {
		fmt.Println("Необходимо указать паттерн для поиска")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	var result []string
	lineCounter := 0
	matchCounter := 0

	beforeBuffer := make([]string, before)

	for scanner.Scan() {
		line := scanner.Text()
		lineCounter++
		matched := matches(line, pattern)

		if invert {
			matched = !matched
		}

		if before > 0 && len(beforeBuffer) > 0 {
			if matched {
				result = append(result, beforeBuffer...)
				beforeBuffer = make([]string, before)
			} else {
				beforeBuffer = append(beforeBuffer[1:], line)
			}
		}

		if matched {
			if count {
				matchCounter++
			} else {
				if lineNum {
					line = fmt.Sprintf("%d:%s", lineCounter, line)
				}
				result = append(result, line)
				for i := 0; i < after && scanner.Scan(); i++ {
					nextLine := scanner.Text()
					if lineNum {
						nextLine = fmt.Sprintf("%d:%s", lineCounter+i+1, nextLine)
					}
					result = append(result, nextLine)
				}
			}
		}
	}

	if count {
		fmt.Println(matchCounter)
	} else {
		for _, line := range result {
			fmt.Println(line)
		}
	}
}

func matches(line, pattern string) bool {
	if ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}
	if fixed {
		return strings.Contains(line, pattern)
	}
	return strings.Contains(line, pattern)
}