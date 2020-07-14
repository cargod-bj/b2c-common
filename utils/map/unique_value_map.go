package mapKits

// Value也是唯一值的map，可用于value查key
type UVMapStrU32 map[string]uint32
type UVMapStrU64 map[string]uint64
type UVMapU32Str map[uint32]string
type UVMapU64Str map[uint64]string

// 根据key获取value
func (n *UVMapStrU32) Key(value uint32) string {
	m := map[string]uint32(*n)
	return GetKeyByValueStrU32(&m, value)
}

// 根据key获取value
func (n *UVMapStrU64) Key(value uint64) string {
	m := map[string]uint64(*n)
	return GetKeyByValueStrU64(&m, value)
}

// 根据key获取value
func (n *UVMapU32Str) Key(value string) uint32 {
	m := map[uint32]string(*n)
	return GetKeyByValueU32Str(&m, value)
}

// 根据key获取value
func (n *UVMapU64Str) Key(value string) uint64 {
	m := map[uint64]string(*n)
	return GetKeyByValueU64Str(&m, value)
}

// 根据key获取value，支持：map[string]uint32
func GetKeyByValueStrU32(m *map[string]uint32, v uint32) string {
	for k := range *m {
		if (*m)[k] == v {
			return k
		}
	}
	return ""
}

// 根据key获取value，支持：map[string]uint64
func GetKeyByValueStrU64(m *map[string]uint64, v uint64) string {
	for k := range *m {
		if (*m)[k] == v {
			return k
		}
	}
	return ""
}

// 根据key获取value，支持：map[uint64]string
func GetKeyByValueU64Str(m *map[uint64]string, v string) uint64 {
	for k := range *m {
		if (*m)[k] == v {
			return k
		}
	}
	return 0
}

// 根据key获取value，支持：map[uint32]string
func GetKeyByValueU32Str(m *map[uint32]string, v string) uint32 {
	for k := range *m {
		if (*m)[k] == v {
			return k
		}
	}
	return 0
}
