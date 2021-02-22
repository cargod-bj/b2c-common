package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"reflect"
	"strconv"
	"time"
)

const (
	timeType         = "time.Time"
	timeTypePtr      = "*time.Time"
	timestampType    = "timestamp.Timestamp"
	timestampTypePtr = "*timestamp.Timestamp"
	uint64Type       = "uint64"
	int64Type        = "int64"
	uint32Type       = "uint32"
	int32Type        = "int32"
	intType          = "intType"
	uintType         = "uintType"
	float64Type      = "float64Type"
	float32Type      = "float32Type"
	decimalTypePtr   = "*decimal.Decimal"
	stringType       = "string"
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
//
//   扩展的跨类型转换，input中的param必须为指针
//   - uint64 to time.Time support second、millisecond、nanosecond to time
//   - uint64 to *time.Time support second、millisecond、nanosecond to time
//   - int64 to time.Time support second、millisecond、nanosecond to time
//   - int64 to *time.Time support second、millisecond、nanosecond to time
//   - *time.Time to uint64 will parse to millisecond
//   - *time.Time to int64 will parse to millisecond
//   - *time.Time to timestamp.Timestamp
//   - *time.Time to *timestamp.Timestamp
//   - *timestamp.Timestamp to time.Time
//   - *timestamp.Timestamp to *time.Time
//   - *decimal.Decimal to string
//   - string to *decimal.Decimal、int、uint、int32、uint32、int64、uint64、float32、float64
//   - time.Time、timestamp.Timestamp and so on, dose not support convert to other type by the soft map
func DecodeDto(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		ErrorUnused:      false,
		DecodeHook: func(inType reflect.Type, outType reflect.Type, src interface{}) (interface{}, error) {

			in := inType.String()
			out := outType.String()

			if in == timeTypePtr {
				srcValue := *src.(*time.Time)
				if int64Type == out {
					return convertTime2Int64(srcValue), nil
				}
				if uint64Type == out {
					var ts int64
					if ts = convertTime2Int64(srcValue); ts < 0 {
						ts = 0
					}
					return uint64(ts), nil
				}
				if timestampType == out {
					result, err := ptypes.TimestampProto(srcValue)
					return &result, err
				}
				if timestampTypePtr == out {
					result, err := ptypes.TimestampProto(srcValue)
					return result, err
				}
			}
			if in == int64Type {
				if timeType == out {
					result := convertInt642Time(src)
					if result == nil {
						return nil, nil
					}
					return *result, nil
				}
				if timeTypePtr == out {
					result := convertInt642Time(src)
					return result, nil
				}
			}
			if in == uint64Type {
				if timeType == out {
					result := convertInt642Time(src)
					if result == nil {
						return nil, nil
					}
					return result, nil
				}
				if timeTypePtr == out {
					result := convertInt642Time(src)
					return result, nil
				}
			}

			if in == timestampTypePtr {
				tmp := src.(*timestamp.Timestamp)
				result, err := ptypes.Timestamp(tmp)
				if timeType == out {
					return result, err
				}
				if timeTypePtr == out {
					return &result, err
				}
			}
			if in == decimalTypePtr {
				temp := src.(*decimal.Decimal)
				if stringType == out {
					result := temp.String()
					return result, nil
				}
			}
			if in == stringType {
				temp := src.(string)
				if decimalTypePtr == out {
					if temp == "" {
						return nil, nil
					}
					result, err := decimal.NewFromString(temp)
					return &result, err
				}
				i, err, done := tryParseNum(temp, out)
				if done {
					return i, err
				}
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

func tryParseNum(temp string, out string) (result interface{}, err error, processed bool) {
	if out == uint64Type || out == int64Type || out == uint32Type || out == int32Type ||
		out == intType || out == uintType || out == float64Type || out == float32Type {
		if temp == "" {
			return 0, nil, true
		}
	} else {
		return nil, nil, false
	}
	switch out {
	case uint64Type:
		result, err = strconv.ParseUint(temp, 10, 64)
		processed = true
	case int64Type:
		result, err = strconv.ParseInt(temp, 10, 64)
		processed = true
	case uint32Type:
		result, err = strconv.ParseUint(temp, 10, 32)
		processed = true
	case int32Type:
		result, err = strconv.ParseInt(temp, 10, 32)
		processed = true
	case intType:
		result, err = strconv.ParseInt(temp, 10, 0)
		processed = true
	case uintType:
		result, err = strconv.ParseUint(temp, 10, 0)
		processed = true
	case float64Type:
		result, err = strconv.ParseFloat(temp, 64)
		processed = true
	case float32Type:
		result, err = strconv.ParseFloat(temp, 32)
		processed = true
	}
	return
}

func convertInt642Time(src interface{}) *time.Time {
	sec, ok := src.(uint64)
	if !ok {
		sec = uint64(src.(int64))
	}
	if sec == 0 {
		return nil
	}
	var result time.Time
	if sec > 1e18 {
		result = time.Unix(0, int64(sec))
	} else if sec > 1e12 {
		s := sec / 1e3
		sec = sec - s*1e3
		if sec < 0 {
			sec = sec + 1e3
			s--
		}
		sec = sec * 1e6
		result = time.Unix(int64(s), int64(sec))
	} else {
		result = time.Unix(int64(sec), 0)
	}
	return &result
}

func convertTime2Int64(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
