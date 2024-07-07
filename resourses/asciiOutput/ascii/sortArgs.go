package ascii

import (
	"fmt"
)

// Methods sorts out the arguments recieved based on the number of arguments written and if the color flag has been passed
func (receive *Receiver) SortArg(filename string, argsPassed []string) string {

	receive.ArgsPassed = argsPassed[0:]

	bannerContent, err := GetFileName(filename)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	receive.FileArr = bannerContent
	return ""
}
