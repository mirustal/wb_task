package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func findAnagramSets(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	result := make(map[string][]string)

	for _, word := range words {
		lowerWord := strings.ToLower(word)    
		sortedWord := sortString(lowerWord)   
		anagrams[sortedWord] = append(anagrams[sortedWord], lowerWord)
	}

	for _, group := range anagrams {
		if len(group) > 1 {
			sort.Strings(group)                
			firstWord := group[0]              
			result[firstWord] = group          
		}
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "друг"}
	anagramSets := findAnagramSets(words)


	for key, set := range anagramSets {
		fmt.Printf("Key: %s, Set: %v\n", key, set)
	}
}
