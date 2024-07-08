package GenerateASCII

import (
	"fmt"

	"github.com/konditi1/ascii-art-web/resourses/asciiOutput/ascii"
)

func GenerateASCII(color, text, filename string) (string, error) {

	var receive ascii.Receiver

	receive.Colorflag = &color
	receive.LettersToColor = ""
	argsPassed := []string{text}

	// Methods sorts-out our arguments to the receive struct
	msg := receive.SortArg(filename, argsPassed)
	if msg != "" {
		return "", fmt.Errorf(msg)
	}
	// Checks if the color flag has been passed
	if color != "" {
		receive.Color()
	} else {
		receive.ColorCode = ""
		receive.Art()
	}
	return receive.Art(), nil
}
