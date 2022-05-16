package reverse_string

func ReverseString(input string) (output string) {
	var a = []rune(input)
	var b []rune
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	output = string(b)
	return output
}
