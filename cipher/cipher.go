package cipher

import (
	"strings"
)

func Encode(keyword string, message string) string {
	encodedMessage := strings.Builder{}
	secret := repeatedSecret(keyword, message)

	for i, r := range message {
		substitution := findSubstitution(rune(secret[i]), r)
		encodedMessage.WriteRune(substitution)
	}

	return encodedMessage.String()
}

func Decode(keyword string, message string) string {
	decodedMessage := strings.Builder{}
	secret := repeatedSecret(keyword, message)

	for i, r := range message {
		row := revertSubstitution(r, rune(secret[i]))
		decodedMessage.WriteRune(row)
	}

	return decodedMessage.String()
}

func Decipher(cipher string, message string) string {
	repeatedSecret := findRepeatedSecret(cipher, message)

	return findSecret(repeatedSecret)
}

func findSecret(repeatedSecret string) string {
	secret := ""

	for i := 1; 2*i < len(repeatedSecret); i++ {
		current := repeatedSecret[:i]
		next := repeatedSecret[i : 2*i]

		if isRepeated(secret, current) {
			// must be repeated
			continue
		}

		if current == next {
			secret = current
		}
	}

	return secret
}
