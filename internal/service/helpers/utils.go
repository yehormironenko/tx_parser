package helpers

import "strconv"

func ConvertHexToInt(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr[2:], 16, 64)
}
