package main

func isalphanumeric1(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isvowel(c byte) bool {
	if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}

func isconsonant(c byte) bool {
	if c >= '0' && c <= '9' {
		return false
	}
	return true
}

// func isdigit(c byte) bool{
//     if (c>='0' && c<='9'){
//         return true
//     }
//     return false
// }

func isValid(word string) bool {
	minchar, vowel, consonant := 0, 0, 0
	for i := 0; i < len(word); i++ {
		if isalphanumeric1(word[i]) {
			// if isdigit(word[i]) {
			//     digit++
			// }else
			if isvowel(word[i]) {
				println("Vowel: ", word[i])
				vowel++
			} else if isconsonant(word[i]) {
				println("Consonant: ", word[i])
				consonant++
			}
			minchar++
		}
	}
	println(minchar, vowel, consonant)
	if minchar > 2 && vowel > 0 && consonant > 0 {
		return true
	}
	return false
}

func main() {
	s := "#zwI"
	println(isValid(s))
}
