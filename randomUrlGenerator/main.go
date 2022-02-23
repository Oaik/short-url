package randomUrlGenerator

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	numOfRandomLetters int = 7
	numOfRandomDigits  int = 3
)

type RandomUrlGenerator struct {
}

func New() *RandomUrlGenerator {
	rand.Seed(time.Now().Unix())
	return &RandomUrlGenerator{}
}

func (e *RandomUrlGenerator) GenerateUrl() string {
	randomUrl := e.generateRandomLetters() + e.generateRandomDigits()
	return randomUrl
}

func (e *RandomUrlGenerator) generateRandomLetters() string {
	randomLettersSlice := make([]string, numOfRandomLetters)
	for i := 0; i < numOfRandomLetters; i++ {
		randomChar := 'a' + rune(rand.Intn(26))
		randomLettersSlice[i] = string(randomChar)
	}
	return strings.Join(randomLettersSlice, "")
}

func (e *RandomUrlGenerator) generateRandomDigits() string {
	randomDigitsSlice := make([]string, numOfRandomDigits)
	for i := 0; i < numOfRandomDigits; i++ {
		randomDigit := rand.Intn(9)
		randomDigitsSlice[i] = strconv.Itoa(randomDigit)
	}
	return strings.Join(randomDigitsSlice, "")
}
