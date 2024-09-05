package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if !strings.HasPrefix(question, "What is ") || !strings.HasSuffix(question, "?") {
		return 0, false
	}

	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimSuffix(question, "?")

	var result int
	words := strings.Fields(question)

	for i, word := range words {
		fmt.Println(i, word)
		prev, err := strconv.Atoi(words[i-1])
		next, err := strconv.Atoi(words[i+1])
		switch word {
		case "plus":
			if err != nil {
				return 0, false
			}
			result = result + prev + next
		case "minus":
			
		}

	}

	return result, true
}
