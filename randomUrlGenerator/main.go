package randomUrlGenerator

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type RandomUrlGenerator struct {
	numOfRandomLetters int
	numOfRandomDigits  int
}

func New(numOfRandomLetters, numOfRandomDigits int) *RandomUrlGenerator {
	rand.Seed(time.Now().Unix())
	e := RandomUrlGenerator{}
	e.numOfRandomLetters = numOfRandomLetters
	e.numOfRandomDigits = numOfRandomDigits
	return &e
}

func (e *RandomUrlGenerator) GenerateUrl() string {
	randomUrl := e.generateRandomLetters() + e.generateRandomDigits()
	return randomUrl
}

func (e *RandomUrlGenerator) generateRandomLetters() string {
	randomLettersSlice := make([]string, e.numOfRandomLetters)
	for i := 0; i < e.numOfRandomLetters; i++ {
		randomChar := 'a' + rune(rand.Intn(26))
		randomLettersSlice[i] = string(randomChar)
	}
	return strings.Join(randomLettersSlice, "")
}

func (e *RandomUrlGenerator) generateRandomDigits() string {
	randomDigitsSlice := make([]string, e.numOfRandomDigits)
	for i := 0; i < e.numOfRandomDigits; i++ {
		randomDigit := rand.Intn(9)
		randomDigitsSlice[i] = strconv.Itoa(randomDigit)
	}
	return strings.Join(randomDigitsSlice, "")
}
