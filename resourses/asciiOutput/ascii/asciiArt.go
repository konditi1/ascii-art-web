package ascii

import (
	"strings"
)

// Ascii prints ASCII art from a given array of characters or writes the result to a file passed by the output.
// The characters are extracted from a predefined file array in the receiver struct.
// The method takes in a receiver stuct and uses the struct fields for forming our output.
// The field wordsArr in the receiver stuct (a slice of strings representing the words to be printed),
// The field ilettersToColor in receiver struct (a string representing the letters to be colored),
// and colordoe filed in the reciver stuct (a string representing the color to be applied).
func Ascii(s Receiver) string {
	var count int
	reset := "\033[0m"
	var outputBuilder strings.Builder

	for _, val := range s.WordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(s.LettersToColor) == 0 && s.ColorCode != "" {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else if strings.Contains(s.LettersToColor, string(v)) {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else {
						outputBuilder.WriteString(s.FileArr[int(start)+i])
					}
				}
				outputBuilder.WriteString("\n")
			}
		} else {
			count++
			if count < len(s.WordsArr) {
				outputBuilder.WriteString("\n")
			}
		}
	}

	return outputBuilder.String()
}
