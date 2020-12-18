package day1

func Part1(expenses []int,sum int) (product int)  {
	product = 0
	expenseMap := make(map[int]int)
	for _, expense := range expenses {
		expenseMap[expense]++
	}
	for key, _ := range expenseMap {
		if expenseMap[sum-key] != 0 {
			product = (sum-key) * key
		}
	}
	return
}
