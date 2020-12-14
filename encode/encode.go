package main

import (
	"fmt"
	"strings"
)

func Splitter(s string) []string {
	words := strings.Fields(s)
	return words
}

func Encode(words []string) string {

	var newString string

	for i, _ := range words {
		w1 := words[i]
		for j := 1; j < len(w1)+1; j++ {
			switch w1[j-1 : j] {
			case "a":
				newString += "1"
			case "e":
				newString += "2"
			case "i":
				newString += "3"
			case "o":
				newString += "4"
			case "u":
				newString += "5"
			default:
				newString += w1[j-1 : j]
			}

		}
		newString += " "
	}

	return newString
}
func main() {
	word := []string(Splitter("hello world"))
	fmt.Printf(Encode(word))

}
