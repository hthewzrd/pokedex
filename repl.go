package main 

import "strings"

func cleanInput(text string) []string {
    word := strings.ToLower(text)

	words := strings.Fields(word)

	return words

}
