package util

// Note:This is generated code, do not edit

func GetIntFrequencyMap(arr []int) map[int]int {
	freqMap := make(map[int]int)
	for _, val := range arr {
		freqMap[val]++
	}
	return freqMap
}