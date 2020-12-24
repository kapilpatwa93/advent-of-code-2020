package day3

func Part1(areaMap [][]string, rightInc, bottomInc int, findObject string) (count int) {
	count = 0
	for i, j := bottomInc, rightInc; i < len(areaMap); i, j = i+bottomInc, (j+rightInc)%len(areaMap[i]) {
		if areaMap[i][j] == findObject {
			count++
		}
	}
	return
}
