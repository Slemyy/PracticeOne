package Recursion

func BruteForcePasswords(password []rune, start int, results *[][]rune) {
	if start == len(password)-1 {
		*results = append(*results, append([]rune(nil), password...))
		return
	}

	for i := start; i < len(password); i++ {
		password[start], password[i] = password[i], password[start]
		BruteForcePasswords(password, start+1, results)
		password[start], password[i] = password[i], password[start]
	}
}
