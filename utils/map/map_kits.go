package mapKits

func KeysStrStr(m *map[string]string) *[]string {
	keys := make([]string, len(*m))
	j := 0
	for k := range *m {
		keys[j] = k
		j++
	}
	return &keys
}

func KeysU32Str(m *map[uint32]string) *[]uint32 {
	keys := make([]uint32, len(*m))
	j := 0
	for k := range *m {
		keys[j] = k
		j++
	}
	return &keys
}

func KeysU64Str(m *map[uint64]string) *[]uint64 {
	keys := make([]uint64, len(*m))
	j := 0
	for k := range *m {
		keys[j] = k
		j++
	}
	return &keys
}

func KeysStrU32(m *map[string]uint32) *[]string {
	keys := make([]string, len(*m))
	j := 0
	for k := range *m {
		keys[j] = k
		j++
	}
	return &keys
}

func KeysStrU64(m *map[string]uint64) *[]string {
	keys := make([]string, len(*m))
	j := 0
	for k := range *m {
		keys[j] = k
		j++
	}
	return &keys
}
