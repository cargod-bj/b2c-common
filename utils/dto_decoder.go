package utils

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

// 使用 mapstructure 解析in中的属性到out中
// 开启了弱匹配，使用如下匹配规则：
//   - bools to string (true = "1", false = "0")
//   - numbers to string (base 10)
//   - bools to int/uint (true = 1, false = 0)
//   - strings to int/uint (base implied by prefix)
//   - int to bool (true if value != 0)
//   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
//     FALSE, false, False. Anything else is an error)
//   - empty array = empty map and vice versa
//   - negative numbers to overflowed uint values (base 10)
//   - slice of maps to a merged map
//   - single values are converted to slices if required. Each
//     element is weakly decoded. For example: "4" can become []int{4}
//     if the target type is an int slice.
//   - uint64 to time.Time
//   - uint64 to *time.Time
//   - time.Time to uint64
//   - *time.Time to uint64
func DecodeDto(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		ErrorUnused:      false,
		DecodeHook: func(inType reflect.Type, outType reflect.Type, src interface{}) (interface{}, error) {
			timeType := "*time.Time"
			timeType2 := "time.Time"
			intType := "uint64"
			if inType.String() == timeType && outType.String() == intType {
				srcValue := src.(*time.Time)
				return uint64(srcValue.Unix() * 1000), nil
			} else if inType.String() == intType && outType.String() == timeType {
				result := time.Unix(int64(src.(uint64)), 0)
				return &result, nil
			} else if inType.String() == timeType2 && outType.String() == intType {
				srcValue := src.(time.Time)
				return uint64(srcValue.Unix()), nil
			} else if inType.String() == intType && outType.String() == timeType2 {
				result := time.Unix(int64(src.(uint64)), 0)
				return result, nil
			}
			return src, nil
		},
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}