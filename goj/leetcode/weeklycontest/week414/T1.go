package week414

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	ymd := strings.Split(date, "-")
	year, _ := strconv.ParseInt(ymd[0], 10, 64)
	month, _ := strconv.ParseInt(ymd[1], 10, 64)
	day, _ := strconv.ParseInt(ymd[2], 10, 64)
	ymd[0], ymd[1], ymd[2] = strconv.FormatInt(year, 2),
		strconv.FormatInt(month, 2),
		strconv.FormatInt(day, 2)
	return strings.Join(ymd, "-")
}
