package utils

import "math/rand"

type NumericTokenGenerator interface {
	isToken(tokens []int, newToken int) bool
	GenerateToken(tokens []int) int
}

type implService struct {
}

func New() NumericTokenGenerator {
	return &implService{}
}

func (impl implService) isToken(tokens []int, newToken int) bool {

	check := false

	for _, token := range tokens {
		if token == newToken {
			check = true
		}
	}

	return check
}

func (impl implService) GenerateToken(tokens []int) int {

	min := 12345
	max := 123456789

	for {
		newToken := rand.Intn(max-min) + min
		check := impl.isToken(tokens, newToken)

		if !check {
			return newToken
		}
	}

}
