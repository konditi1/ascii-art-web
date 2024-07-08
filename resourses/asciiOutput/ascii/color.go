package ascii

// Color takes words a slice of string from a certain index, joins the words,
// parameters: colorflag string, lettersTocolor string, argsPassed []string, bannerContent []string
// and calls the Ascii function to print the words in ascii-art
func (s *Receiver) Color() {

	if len(s.ArgsPassed) == 1 {
		s.IndexToStartDisplay = 0
	}
	s.Art()
}
