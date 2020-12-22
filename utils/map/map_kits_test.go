package mapKits

import "testing"

func TestKeysStrU32(t *testing.T) {
	m := make(map[string]uint32)
	m["a"] = 1
	m["b"] = 1
	m["c"] = 1
	keys := KeysStrU32(&m)
	for i := range *keys {
		println("KeysStrU32 -> ", (*keys)[i])
	}
	t.Fatal("测试失败")
}

func TestMap(t *testing.T) {
	m := make(map[string]uint32)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8
	for k, v := range m {
		println("KeysStrU32 -> ", v)
		if v%2 == 0 {
			delete(m, k)
		}
	}
	println("------------------------")
	for _, v := range m {
		println("KeysStrU32 -> ", v)
	}
	t.Fatal("测试失败")
}
