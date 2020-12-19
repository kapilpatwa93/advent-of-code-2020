package day1

func Part2(expenses []int, sum int) (product int) {
	product = 0
	expenseMap := make(map[int]int)
	for _, expense := range expenses {
		expenseMap[expense]++
	}
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenseMap[sum-(expenses[i]+expenses[j])] != 0 {
				product = expenses[i] * expenses[j] * (sum - (expenses[i] + expenses[j]))
				return
			}
		}
	}
	return
}
