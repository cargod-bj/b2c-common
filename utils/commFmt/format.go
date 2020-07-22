package commFmt

import (
	"fmt"
	"math"
	"strconv"
)

// float64保留n位有效数字转string
func Formatf(value float64) string {
	result := fmt.Sprintf("%.2f", value)
	return result
}

// 主要逻辑就是先乘，trunc之后再除回去，就达到了保留N位小数的效果
func FormatFloat(num float64, n int) string {
	// 默认乘1
	d := float64(1)
	if n > 0 {
		// 10的N次方
		d = math.Pow10(n)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
}
