package ascii

// Art takes words a slice of string from a certain index, joins the words,
// checks if there is a non-ascii character
// and calls the Ascii method which writes the result to a specified file or ot the terminal
func (s *Receiver) Art() {
	word := Arrange(s.ArgsPassed[s.IndexToStartDisplay:])
	s.WordsArr = Slice(word)
	if !CheckAscii(s.WordsArr) {
		return
	}
	Ascii(*s)
}
