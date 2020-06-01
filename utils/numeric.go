package utils

// uint32的最小值
const UINT32_MIN uint32 = 0

// uint32的最大值
const UINT32_MAX uint32 = ^uint32(0)

// int32最大值
const INT32_MAX int32 = int32(UINT32_MAX >> 1)

// int32最小值
const INT32_MIN int32 = ^INT32_MAX

// uint64最大值
const UINT64_MAX uint64 = ^uint64(0)

// uint64最小值
const UINT64_MIN uint64 = 0

// int64最大值
const INT64_MAX int64 = int64(UINT64_MAX >> 1)

// int64最小值
const INT64_MIN int64 = ^INT64_MAX

// uint最大值
const UINT_MAX uint = ^uint(0)

// uint最小值
const UINT_MIN uint = 0

// int最大值
const INT_MAX int = int(UINT_MAX >> 1)

// int最小值
const INT_MIN int = ^INT_MAX
