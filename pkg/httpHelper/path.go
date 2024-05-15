package httpHelper

import (
	"strconv"
)

func GetIdFromString(id string) int {
	if id == "" {
		return -1
	}
	if id[0] == ('0') {
		return -1
	}
	num, err := strconv.Atoi(id)
	if err != nil {
		return -1
	}
	return num
}
