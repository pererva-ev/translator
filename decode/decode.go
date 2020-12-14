package main

import (
	"fmt"
	"strings"
)

func Splitter(s string) []string {
	words := strings.Fields(s)
	return words
}

func Decode(words []string) string {

	var newString string

	for i, _ := range words {
		w1 := words[i]
		for j := 1; j < len(w1)+1; j++ {
			switch w1[j-1 : j] {
			case "1":
				newString += "a"
			case "2":
				newString += "e"
			case "3":
				newString += "i"
			case "4":
				newString += "o"
			case "5":
				newString += "u"
			default:
				newString += w1[j-1 : j]
			}
		}
		newString += " "
	}

	return newString
}
func main() {
	word := []string(Splitter("h2ll4 w4rld"))
	fmt.Printf(Decode(word))

}
