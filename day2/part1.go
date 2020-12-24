package day2

import (
	"github.com/kapilpatwa93/advent-of-code-2020/util"
	"strings"
)

type Policy1 struct {
	LeastCount int
	MostCount  int
	Password   string
	Character  string
}

func Part1(policies []Policy1) (validPasswordCount int) {
	validPasswordCount = 0
	for _, policy := range policies {
		if policy.IsValidPassword() {
			validPasswordCount++
		}
	}
	return
}

func (policy Policy1) IsValidPassword() (isValid bool) {
	isValid = false
	passWordFreqMap := util.GetStringFrequencyMap(strings.Split(policy.Password, ""))
	if passWordFreqMap[policy.Character] >= policy.LeastCount && passWordFreqMap[policy.Character] <= policy.MostCount {
		isValid = true
	}
	return
}
