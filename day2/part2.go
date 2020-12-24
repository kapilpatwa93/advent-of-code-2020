package day2

type Policy2 struct {
	Position1 int
	Position2 int
	Password  string
	Character string
}

func Part2(policies []Policy2) (validPasswordCount int) {
	validPasswordCount = 0
	for _, policy := range policies {
		if policy.IsValidPassword() {
			validPasswordCount++
		}
	}
	return
}

func (policy Policy2) IsValidPassword() (isValid bool) {
	isValid = false
	matchCount := 0
	if policy.Password[policy.Position1-1:policy.Position1] == policy.Character {
		matchCount++
	}
	if policy.Password[policy.Position2-1:policy.Position2] == policy.Character {
		matchCount++
	}
	if matchCount == 1 {
		isValid = true
	}
	return
}
