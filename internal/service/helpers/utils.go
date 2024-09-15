package helpers

import (
	"fmt"
	"strconv"
)

func ConvertHexToInt(hexStr string) (int64, error) {
	if len(hexStr) < 3 || hexStr[:2] != "0x" {
		return 0, fmt.Errorf("invalid hexadecimal string: %s", hexStr)
	}

	return strconv.ParseInt(hexStr[2:], 16, 64)
}
