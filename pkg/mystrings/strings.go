package mystrings

func HasAllUniqueRunes(input string) bool {
	chars := make(map[rune]bool)
	for _, i := range input {
		if _, ok := chars[i]; ok {
			return false
		}
		chars[i] = true
	}
	return true
}

func IsPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
