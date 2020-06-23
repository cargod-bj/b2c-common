package sliceKits

import "testing"

func TestAdd(t *testing.T) {
	uint32Slice := []uint32{1, 2, 3, 4}
	uint32Slice = GetU32(uint32Slice).Add(5)
	for i := range uint32Slice {
		println("uint32Slice -> %d", uint32Slice[i])
	}
	t.Fatal("测试失败")
}
