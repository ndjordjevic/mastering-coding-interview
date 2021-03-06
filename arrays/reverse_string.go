package arrays

// Tijana -> anajiT
// Nenad -> daneN
func reverseString(str string) string {
	runes := []rune(str)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

// O(n), O(n)
