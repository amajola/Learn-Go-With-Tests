package iteration

func Repeat(character string) string {
	var reapeated string
	for i := 0; i < 5; i++ {
		reapeated += character
	}
	return reapeated
}
