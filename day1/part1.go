package day1

import "github.com/kapilpatwa93/advent-of-code-2020/util"

func Part1(expenses []int,sum int) (product int)  {
	product = 0
	expenseMap := util.GetIntFrequencyMap(expenses)
	for key, _ := range expenseMap {
		if expenseMap[sum-key] != 0 {
			product = (sum-key) * key
		}
	}
	return
}
