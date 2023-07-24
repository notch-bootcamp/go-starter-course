package returns

func SplitInHalf(word string) (string, string) {
	half := len(word) / 2
	return word[:half], word[half:]
}
