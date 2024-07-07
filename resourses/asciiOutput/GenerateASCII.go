package GenerateASCII

import (
	"fmt"

	"github.com/konditi1/ascii-art-web/resourses/asciiOutput/ascii"
)

func GenerateASCII(color, text, filename string) {

	var receive ascii.Receiver

	receive.Colorflag = &color
	receive.LettersToColor = ""
	argsPassed := []string{text}

	// Methods sorts-out our arguments to the receive struct
	msg := receive.SortArg(filename, argsPassed)
	if msg != "" {
		fmt.Println(msg)
		return
	}
	// Checks if the color flag has been passed
	if color != "" {
		receive.Color()
	} else {
		receive.ColorCode = ""
		receive.Art()
	}
}
