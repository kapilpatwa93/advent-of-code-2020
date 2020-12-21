package util

// Note:This is generated code, do not edit

func GetFloat32FrequencyMap(arr []float32) map[float32]int {
	freqMap := make(map[float32]int)
	for _, val := range arr {
		freqMap[val]++
	}
	return freqMap
}