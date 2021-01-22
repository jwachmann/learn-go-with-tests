package iteration

// Repeat creates and returns a string which is the given character repeated 5 times
func Repeat(character string, repeatCount int) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
