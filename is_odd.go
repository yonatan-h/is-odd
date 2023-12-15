package is_odd

import "github.com/yonatan-h/is_odd/utils"


func IsOdd(num int) bool {
	isOdd := utils.Equal(utils.Remainder(num, 2), 1)
	return isOdd
}