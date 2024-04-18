package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fields string
	var delimiter string
	var separated bool

	flag.StringVar(&fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "только строки с разделителем")

	flag.Parse()

	if fields == "" {
		log.Fatal("Не указаны колонки для вывода через параметр -f")
	}


	fieldIndexes := parseFields(fields)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if separated && !strings.Contains(txt, delimiter) {
			continue
		}

		splitTxt := strings.Split(txt, delimiter)
		output := selectFields(splitTxt, fieldIndexes)
		fmt.Println(strings.Join(output, delimiter))
	}
}


func parseFields(fieldStr string) []int {
	fieldStrs := strings.Split(fieldStr, ",")
	var fields []int
	for _, field := range fieldStrs {
		var index int
		fmt.Sscanf(field, "%d", &index)
		if index > 0 { 
			fields = append(fields, index-1) 
		}
	}
	return fields
}


func selectFields(record []string, fields []int) []string {
	var result []string
	for _, fieldIndex := range fields {
		if fieldIndex < len(record) {
			result = append(result, record[fieldIndex])
		}
	}
	return result
}