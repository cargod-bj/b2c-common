package mapKits

import "testing"

func TestUVMapStrU32(t *testing.T) {
	m := make(map[string]uint32)
	keys := KeysStrU32(&m)
	for i := range *keys {
		println("KeysStrU32 -> %s", (*keys)[i])
	}
	t.Fatal("测试失败")
}
