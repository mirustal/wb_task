package main
import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	column    int
	numeric   bool
	reverse   bool
	unique    bool
	inputFile string
	outputFile string
)

func init() {
	flag.IntVar(&column, "k", 0, "Колонка для сортировки")
	flag.BoolVar(&numeric, "n", false, "Числовая сортировка")
	flag.BoolVar(&reverse, "r", false, "Сортировать в ону братном порядке")
	flag.BoolVar(&unique, "u", false, "Не выводить повторяющиеся строки")
	flag.StringVar(&inputFile, "i", "", "Входной файл")
	flag.StringVar(&outputFile, "o", "", "Выходной файл")
	flag.Parse()
}

func main() {
	if inputFile == "" || outputFile == "" {
		fmt.Println("Необходимо указать входной и выходной файлы")
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.SliceStable(lines, func(i, j int) bool {
		iCols := strings.Fields(lines[i])
		jCols := strings.Fields(lines[j])

		var iKey, jKey string
		if column < len(iCols) {
			iKey = iCols[column]
		}
		if column < len(jCols) {
			jKey = jCols[column]
		}

		if numeric {
			iVal, err1 := strconv.Atoi(iKey)
			jVal, err2 := strconv.Atoi(jKey)
			if err1 == nil && err2 == nil {
				if reverse {
					return iVal > jVal
				}
				return iVal < jVal
			}
		}

		if reverse {
			return iKey > jKey
		}
		return iKey < jKey
	})

	if unique {
		lines = uniqueLines(lines)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}

func uniqueLines(lines []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range lines {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}