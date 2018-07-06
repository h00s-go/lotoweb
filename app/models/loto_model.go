package models

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// LotoModel for Loto methods
type LotoModel struct {
}

// NewLotoModel creates new loto model
func NewLotoModel() *LotoModel {
	return &LotoModel{}
}

// GetNumbers returns numbers
func (lm *LotoModel) GetNumbers(count int, max int) string {
	var numbers []int
	rand.Seed(time.Now().UnixNano())

	for len(numbers) < count {
		n := rand.Intn(max) + 1
		if !contains(numbers, n) {
			numbers = append(numbers, n)
		}
	}

	sort(numbers)
	var numbersAsText []string
	for _, i := range numbers {
		numbersAsText = append(numbersAsText, strconv.Itoa(i))
	}

	return strings.Join(numbersAsText, ", ")
}

func contains(s []int, v int) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

func sort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
