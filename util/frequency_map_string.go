package util

// Note:This is generated code, do not edit

func GetStringFrequencyMap(arr []string) map[string]int {
	freqMap := make(map[string]int)
	for _, val := range arr {
		freqMap[val]++
	}
	return freqMap
}