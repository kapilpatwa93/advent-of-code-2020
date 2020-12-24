package day3

func Part2(areaMap [][]string, incrCoordinates [][]int, findObject string) (product int) {
	product = 1
	for _, coordinate := range incrCoordinates {
		product *= Part1(areaMap, coordinate[0], coordinate[1], findObject)
	}
	return
}
