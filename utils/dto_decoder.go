package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
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
//   - uint64 to time.Time support second、millisecond、nanosecond to time
//   - uint64 to *time.Time support second、millisecond、nanosecond to time
//   - *time.Time to uint64 will parse to millisecond
//   - int64 to time.Time support second、millisecond、nanosecond to time
//   - int64 to *time.Time support second、millisecond、nanosecond to time
//   - *time.Time to int64 will parse to millisecond
//   - *timestamp.Timestamp to time.Time
//   - *time.Time to timestamp.Timestamp
//   - *decimal.Decimal to string
//   - string to *decimal.Decimal
//   - time.Time、timestamp.Timestamp
func DecodeDto(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		ErrorUnused:      false,
		DecodeHook: func(inType reflect.Type, outType reflect.Type, src interface{}) (interface{}, error) {
			timeType := "time.Time"
			timeTypePtr := "*time.Time"
			timestampType := "timestamp.Timestamp"
			timestampTypePtr := "*timestamp.Timestamp"
			uint64Type := "uint64"
			int64Type := "int64"
			decimalTypePtr := "*decimal.Decimal"
			stringType := "string"

			in := inType.String()
			out := outType.String()

			if in == timeType {
				srcValue := src.(time.Time)
				if int64Type == out {
					return int64(convertTime2Uint64(srcValue)), nil
				}
				if uint64Type == out {
					return convertTime2Uint64(srcValue), nil
				}
				if timestampType == out {
					tmp, err := ptypes.TimestampProto(srcValue)
					if err == nil {
						return *tmp, err
					}
					return nil, err
				}
				if timestampTypePtr == out {
					result, err := ptypes.TimestampProto(srcValue)
					return result, err
				}
			}
			if in == timeTypePtr {
				srcValue := *src.(*time.Time)
				if int64Type == out {
					return int64(convertTime2Uint64(srcValue)), nil
				}
				if uint64Type == out {
					return convertTime2Uint64(srcValue), nil
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
					return result, nil
				}
				if timeTypePtr == out {
					result := convertInt642Time(src)
					return &result, nil
				}
			}
			if in == uint64Type {
				if timeType == out {
					result := convertInt642Time(src)
					return result, nil
				}
				if timeTypePtr == out {
					result := convertInt642Time(src)
					return &result, nil
				}
			}
			if in == timestampType {
				tmp := src.(timestamp.Timestamp)
				result, err := ptypes.Timestamp(&tmp)
				if timeType == out {
					return result, err
				}
				if timeTypePtr == out {
					return &result, err
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
					result, err := decimal.NewFromString(temp)
					return &result, err
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

func convertInt642Time(src interface{}) time.Time {
	sec, ok := src.(uint64)
	if !ok {
		sec = uint64(src.(int64))
	}
	if sec > 1e18 {
		return time.Unix(0, int64(sec))
	} else if sec > 1e12 {
		s := sec / 1e3
		sec = sec - s*1e3
		if sec < 0 {
			sec = sec + 1e3
			s--
		}
		sec = sec * 1e6
		return time.Unix(int64(s), int64(sec))
	}
	return time.Unix(int64(sec), 0)
}

func convertTime2Uint64(t time.Time) uint64 {
	return uint64(t.UnixNano() / 1e9)
}

/**
 *时间和字符串转换，需要传入字符串格式化类型format
 *此方法只能支持一种类型的时间转换，由于只传了一种format
 */
func DecodeStringDto(input, output interface{}, format string) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		ErrorUnused:      false,
		DecodeHook: func(inType reflect.Type, outType reflect.Type, src interface{}) (interface{}, error) {
			timeType := "*time.Time"
			timeTypePrt := "time.Time"
			timestampType := "timestamp.Timestamp"
			timestampTypePtr := "*timestamp.Timestamp"
			int64Type := "int64"
			intType := "uint64"
			stringType := "string"
			if inType.String() == stringType && outType.String() == intType {
				srcValue, err := time.Parse(format, src.(string))
				return uint64(srcValue.Unix() * 1000), err
			} else if inType.String() == intType && outType.String() == stringType {
				res := time.Unix(int64(src.(uint64)), 0)
				result := res.Format(format)
				return &result, nil
			} else if inType.String() == timeTypePrt && outType.String() == stringType {
				srcValue := src.(time.Time)
				return srcValue.Format(format), nil
			} else if inType.String() == stringType && outType.String() == timeTypePrt {
				result, err := time.Parse(format, src.(string))
				return result, err
			} else if inType.String() == timeType && outType.String() == stringType {
				srcValue := src.(*time.Time)
				return srcValue.Format(format), nil
			} else if inType.String() == stringType && outType.String() == timeType {
				result, err := time.Parse(format, src.(string))
				return &result, err
			} else if inType.String() == stringType && outType.String() == int64Type {
				srcValue, err := time.Parse(format, src.(string))
				return int64(srcValue.Unix() * 1000), err
			} else if inType.String() == int64Type && outType.String() == stringType {
				result := time.Unix(src.(int64), 0)
				return result.Format(format), nil
			} else if inType.String() == timestampTypePtr && outType.String() == stringType {
				result, err := ptypes.Timestamp(src.(*timestamp.Timestamp))
				return result.Format(format), err
			} else if inType.String() == stringType && outType.String() == timestampTypePtr {
				temp, err := time.Parse(format, src.(string))
				if err != nil {
					return nil, err
				}
				result, err := ptypes.TimestampProto(temp)
				return result, err
			} else if inType.String() == timestampType && outType.String() == stringType {
				tmp := src.(timestamp.Timestamp)
				result, err := ptypes.Timestamp(&tmp)
				return result.Format(format), err
			} else if inType.String() == stringType && outType.String() == timestampType {
				temp, err := time.Parse(format, src.(string))
				if err != nil {
					return nil, err
				}
				result, err := ptypes.TimestampProto(temp)
				return &result, err
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
