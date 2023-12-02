package util

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		tmp := runes[i]
		runes[i] = runes[j]
		runes[j] = tmp
	}
	return string(runes)
}
