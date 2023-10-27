package utils

import (
	"math/rand"
)

var symbolsRunes = []rune("zxcvbnmasdfghjklqwertyuiop123456789")

func GenerateUID() string {
	b := make([]rune, 19)
	for i := range b {
		b[i] = symbolsRunes[rand.Intn(len(symbolsRunes))]
	}
	return string(b)
}
