package cipher

import (
	"math"
	"strings"
)

const numAlphabets = 'z' - 'a' + 1

func findSubstitution(column rune, row rune) rune {
	substitution := column + row - 'a'

	if substitution > 'z' {
		return substitution - numAlphabets
	}

	return substitution
}

func revertSubstitution(substitution rune, r rune) rune {
	out := substitution - r + 'a'

	if out < 'a' {
		return out + numAlphabets
	}

	return out
}

func repeatedSecret(secret string, message string) string {
	repeated := math.Round(float64(len(message)/len(secret)) + 0.5)

	return strings.Repeat(secret, int(repeated))[:len(message)]
}

func findRepeatedSecret(cipher string, message string) string {
	repeatedSecret := strings.Builder{}

	for i, s := range cipher {
		column := revertSubstitution(s, rune(message[i]))
		repeatedSecret.WriteRune(column)
	}

	return repeatedSecret.String()
}

func isRepeated(secret string, message string) bool {
	return len(secret) > 0 && len(message)%len(secret) == 0 && strings.Repeat(secret, len(message)/len(secret)) == message
}
