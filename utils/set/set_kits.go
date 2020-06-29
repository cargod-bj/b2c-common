package setKits

type setU64 map[uint64]struct{}

func SetU64() setU64 {
	return make(map[uint64]struct{})
}

type ISetU64 interface {
	Add(e uint64)
	Delete(e uint64)
	Contain(e uint64) bool
	Slice() *[]uint64
}

func (s *setU64) Add(e uint64) {
	(*s)[e] = struct{}{}
}

func (s *setU64) Delete(e uint64) {
	delete(*s, e)
}

func (s *setU64) Contain(e uint64) bool {
	_, ok := (*s)[e]
	return ok
}

func (s *setU64) Slice() *[]uint64 {
	keys := make([]uint64, len(*s))
	j := 0
	for k := range *s {
		keys[j] = k
		j++
	}
	return &keys
}

type SetU32 map[uint32]struct{}

type ISetU32 interface {
	Add(e uint32)
	Delete(e uint32)
	Contain(e uint32) bool
	Slice() *[]uint32
}

func (s *SetU32) Add(e uint32) {
	(*s)[e] = struct{}{}
}

func (s *SetU32) Delete(e uint32) {
	delete(*s, e)
}

func (s *SetU32) Contain(e uint32) bool {
	_, ok := (*s)[e]
	return ok
}

func (s *SetU32) Slice() *[]uint32 {
	keys := make([]uint32, len(*s))
	j := 0
	for k := range *s {
		keys[j] = k
		j++
	}
	return &keys
}

type SetStr map[string]struct{}

type ISetStr interface {
	Add(e string)
	Delete(e string)
	Contain(e string) bool
	Slice() *[]string
}

func (s *SetStr) Add(e string) {
	(*s)[e] = struct{}{}
}

func (s *SetStr) Delete(e string) {
	delete(*s, e)
}

func (s *SetStr) Contain(e string) bool {
	_, ok := (*s)[e]
	return ok
}

func (s *SetStr) Slice() *[]string {
	keys := make([]string, len(*s))
	j := 0
	for k := range *s {
		keys[j] = k
		j++
	}
	return &keys
}
