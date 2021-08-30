package helpers

import "strconv"

func ParseStringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return result
}

func Contains(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func GetOffset(page int, limit int) int {
	offset := (page - 1) * limit
	if offset < 0 {
		return 0
	}
	return offset
}
