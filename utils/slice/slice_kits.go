package sliceKits

type uint32Slice []uint32
type uint64Slice []uint64
type strSlice []string

type SliceUint32Kits interface {
	IndexOf(e uint32) int
	Add(e uint32) uint32Slice
	AddIfAbsent(e uint32) uint32Slice
	Remove(e uint32) uint32Slice
	Replace(old uint32, new uint32) uint32Slice
	Contain(e uint32) bool
}

type SliceUint64Kits interface {
	IndexOf(e uint64) int
	Add(e uint64) uint64Slice
	AddIfAbsent(e uint64) uint64Slice
	Remove(e uint64) uint64Slice
	Replace(old uint64, new uint64) uint64Slice
	Contain(e uint64) bool
}

type SliceStrKits interface {
	IndexOf(e string) int
	Add(e string) strSlice
	AddIfAbsent(e string) strSlice
	Remove(e string) strSlice
	Replace(old string, new string) strSlice
	Contain(e string) bool
}

func GetU32(original []uint32) uint32Slice {
	return original
}

func GetU64(original []uint64) uint64Slice {
	return original
}

func GetStr(original []string) strSlice {
	return original
}

func (s uint32Slice) IndexOf(e uint32) int {
	if len(s) == 0 {
		return -1
	}
	for i, _ := range s {
		if (s)[i] == e {
			return i
		}
	}
	return -1
}

func (s uint32Slice) Add(e uint32) uint32Slice {
	return append(s, e)
}

func (s uint32Slice) AddIfAbsent(e uint32) uint32Slice {
	index := s.IndexOf(e)
	if index != -1 {
		return s
	}
	return s.Add(e)
}

func (s uint32Slice) Remove(e uint32) uint32Slice {
	index := s.IndexOf(e)
	if index != -1 {
		return append(s[:index], s[index:]...)
	}
	return s
}

func (s uint32Slice) Replace(old uint32, new uint32) uint32Slice {
	for i := range s {
		if s[i] == old {
			s[i] = new
		}
	}
	return s
}

func (s uint32Slice) Contain(e uint32) bool {
	return s.IndexOf(e) != -1
}

func (s uint64Slice) IndexOf(e uint64) int {
	if len(s) == 0 {
		return -1
	}
	for i, _ := range s {
		if (s)[i] == e {
			return i
		}
	}
	return -1
}

func (s uint64Slice) Add(e uint64) uint64Slice {
	return append(s, e)
}

func (s uint64Slice) AddIfAbsent(e uint64) uint64Slice {
	index := s.IndexOf(e)
	if index != -1 {
		return s
	}
	return s.Add(e)
}

func (s uint64Slice) Remove(e uint64) uint64Slice {
	index := s.IndexOf(e)
	if index != -1 {
		return append(s[:index], s[index:]...)
	}
	return s
}

func (s uint64Slice) Replace(old uint64, new uint64) uint64Slice {
	for i := range s {
		if s[i] == old {
			s[i] = new
		}
	}
	return s
}

func (s uint64Slice) Contain(e uint64) bool {
	return s.IndexOf(e) != -1
}

func (s strSlice) IndexOf(e string) int {
	if len(s) == 0 {
		return -1
	}
	for i, _ := range s {
		if (s)[i] == e {
			return i
		}
	}
	return -1
}

func (s strSlice) Add(e string) strSlice {
	return append(s, e)
}

func (s strSlice) AddIfAbsent(e string) strSlice {
	index := s.IndexOf(e)
	if index != -1 {
		return s
	}
	return s.Add(e)
}

func (s strSlice) Remove(e string) strSlice {
	index := s.IndexOf(e)
	if index != -1 {
		return append(s[:index], s[index:]...)
	}
	return s
}

func (s strSlice) Replace(old string, new string) strSlice {
	for i := range s {
		if s[i] == old {
			s[i] = new
		}
	}
	return s
}

func (s strSlice) Contain(e string) bool {
	return s.IndexOf(e) != -1
}
