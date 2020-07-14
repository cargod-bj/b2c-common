package mapKits

type uniqueValueMapStrU32Impl struct {
	o *map[string]uint32
	b map[uint32]string
}

func (n *uniqueValueMapStrU32Impl) GetKey(value uint32) string {
	if n.b == nil {
		return GetKeyByValueStrU32(n.o, value)
	}
	return n.b[value]
}

type UniqueValueMapStrU32 interface {
	GetKey(value uint32) string
}

// 获取Value值唯一的，可value查key，反向查询的map对象
func UVMapStrU32(m *map[string]uint32) UniqueValueMapStrU32 {
	impl := uniqueValueMapStrU32Impl{o: m}
	if len(*m) < 100000 {
		b := map[uint32]string{}
		for k := range *m {
			v := (*m)[k]
			b[v] = k
		}
		impl.b = b
	}
	return &impl
}

type uniqueValueMapStrU64Impl struct {
	o *map[string]uint64
	b map[uint64]string
}

func (n *uniqueValueMapStrU64Impl) GetKey(value uint64) string {
	if n.b == nil {
		return GetKeyByValueStrU64(n.o, value)
	}
	return n.b[value]
}

type UniqueValueMapStrU64 interface {
	GetKey(value uint64) string
}

// 获取Value值唯一的，可value查key，反向查询的map对象
func UVMapStrU64(m *map[string]uint64) UniqueValueMapStrU64 {
	impl := uniqueValueMapStrU64Impl{o: m}
	if len(*m) < 100000 {
		b := map[uint64]string{}
		for k := range *m {
			v := (*m)[k]
			b[v] = k
		}
		impl.b = b
	}
	return &impl
}

type uniqueValueMapU32StrImpl struct {
	o *map[uint32]string
	b map[string]uint32
}

func (n *uniqueValueMapU32StrImpl) GetKey(value string) uint32 {
	if n.b == nil {
		return GetKeyByValueU32Str(n.o, value)
	}
	return n.b[value]
}

type UniqueValueMapU32Str interface {
	GetKey(value string) uint32
}

// 获取Value值唯一的，可value查key，反向查询的map对象
func UVMapU32Str(m *map[uint32]string) UniqueValueMapU32Str {
	impl := uniqueValueMapU32StrImpl{o: m}
	if len(*m) < 100000 {
		b := map[string]uint32{}
		for k := range *m {
			v := (*m)[k]
			b[v] = k
		}
		impl.b = b
	}
	return &impl
}

type uniqueValueMapU64StrImpl struct {
	o *map[uint64]string
	b map[string]uint64
}

func (n *uniqueValueMapU64StrImpl) GetKey(value string) uint64 {
	if n.b == nil {
		return GetKeyByValueU64Str(n.o, value)
	}
	return n.b[value]
}

type UniqueValueMapU64Str interface {
	GetKey(value string) uint64
}

// 获取Value值唯一的，可value查key，反向查询的map对象
func UVMapU64Str(m *map[uint64]string) UniqueValueMapU64Str {
	impl := uniqueValueMapU64StrImpl{o: m}
	if len(*m) < 100000 {
		b := map[string]uint64{}
		for k := range *m {
			v := (*m)[k]
			b[v] = k
		}
		impl.b = b
	}
	return &impl
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
