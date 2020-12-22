package goroutineKits

import (
	"bytes"
	"runtime"
	"strconv"
)

// 获取当前goroutineID
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	println("currentId:", n)
	return n
}
