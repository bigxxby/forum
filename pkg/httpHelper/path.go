package httpHelper

import (
	"regexp"
	"strconv"
)

func GetPostIdFromPath(path string) int {
	re := regexp.MustCompile(`/api/posts/(\w+)`)
	matches := re.FindStringSubmatch(path)
	if len(matches) != 2 {
		return -1
	}
	if matches[1][0] == ('0') {
		return -1
	}
	num, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return -1
	}
	return int(num)
}
