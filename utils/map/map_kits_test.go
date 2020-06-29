package mapKits

import "testing"

func TestKeysStrU32(t *testing.T) {
	m := make(map[string]uint32)
	m["a"] = 1
	m["b"] = 1
	m["c"] = 1
	keys := KeysStrU32(&m)
	for i := range *keys {
		println("KeysStrU32 -> %s", (*keys)[i])
	}
	t.Fatal("测试失败")
}
