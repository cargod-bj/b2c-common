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
			timeTypePrt := "time.Time"
			timestampType := "timestamp.Timestamp"
			timestampTypePtr := "*timestamp.Timestamp"
			int64Type := "int64"
			uint64Type := "uint64"
			decimalTypePtr := "*decimal.Decimal"
			stringType := "string"
			if inType.String() == timeType && outType.String() == uint64Type {
				srcValue := src.(*time.Time)
				return uint64(srcValue.Unix() * 1000), nil
			} else if inType.String() == uint64Type && outType.String() == timeType {
				result := time.Unix(int64(src.(uint64)), 0)
				return &result, nil
			} else if inType.String() == timeTypePrt && outType.String() == uint64Type {
				srcValue := src.(time.Time)
				return uint64(srcValue.Unix()), nil
			} else if inType.String() == uint64Type && outType.String() == timeTypePrt {
				result := time.Unix(int64(src.(uint64)), 0)
				return result, nil
			} else if inType.String() == timeType && outType.String() == int64Type {
				srcValue := src.(*time.Time)
				return int64(srcValue.Unix() * 1000), nil
			} else if inType.String() == int64Type && outType.String() == timeType {
				result := time.Unix(src.(int64), 0)
				return &result, nil
			} else if inType.String() == timeTypePrt && outType.String() == int64Type {
				srcValue := src.(time.Time)
				return int64(srcValue.Unix()), nil
			} else if inType.String() == int64Type && outType.String() == timeTypePrt {
				result := time.Unix(src.(int64), 0)
				return result, nil
			} else if inType.String() == timestampTypePtr && outType.String() == timeType {
				result, err := ptypes.Timestamp(src.(*timestamp.Timestamp))
				return &result, err
			} else if inType.String() == timestampTypePtr && outType.String() == timeTypePrt {
				result, err := ptypes.Timestamp(src.(*timestamp.Timestamp))
				return result, err
			} else if (inType.String() == timeType || inType.String() == timeTypePrt) && outType.String() == timestampTypePtr {
				result, err := ptypes.TimestampProto(src.(time.Time))
				return result, err
			} else if inType.String() == timestampType && outType.String() == timeType {
				tmp := src.(timestamp.Timestamp)
				result, err := ptypes.Timestamp(&tmp)
				return result, err
			} else if inType.String() == timestampType && outType.String() == timeTypePrt {
				tmp := src.(timestamp.Timestamp)
				result, err := ptypes.Timestamp(&tmp)
				return &result, err
			} else if (inType.String() == timeType || inType.String() == timeTypePrt) && outType.String() == timestampType {
				result, err := ptypes.TimestampProto(src.(time.Time))
				return &result, err
			} else if (inType.String() == stringType) && outType.String() == decimalTypePtr {
				result, err := decimal.NewFromString(src.(string))
				return &result, err
			} else if (inType.String() == decimalTypePtr) && outType.String() == stringType {
				result := (src.(*decimal.Decimal)).String()
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
			uint64Type := "uint64"
			uint32Type := "uint32"
			int32Type := "int32"
			intType := "int"
			stringType := "string"
			if inType.String() == stringType && outType.String() == uint64Type {
				srcValue, err := time.Parse(format, src.(string))
				return uint64(srcValue.Unix() * 1000), err
			} else if inType.String() == uint64Type && outType.String() == stringType {
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
			} else if inType.String() == stringType && outType.String() == uint64Type {
				temp, err := strconv.Atoi(src.(string))
				if err != nil {
					return nil, err
				}
				result := uint64(temp)
				return &result, err
			} else if inType.String() == stringType && outType.String() == int64Type {
				temp, err := strconv.Atoi(src.(string))
				if err != nil {
					return nil, err
				}
				result := int64(temp)
				return &result, err
			} else if inType.String() == stringType && outType.String() == uint32Type {
				temp, err := strconv.Atoi(src.(string))
				if err != nil {
					return nil, err
				}
				result := uint32(temp)
				return &result, err
			} else if inType.String() == stringType && outType.String() == int32Type {
				temp, err := strconv.Atoi(src.(string))
				if err != nil {
					return nil, err
				}
				result := int32(temp)
				return &result, err
			} else if inType.String() == stringType && outType.String() == intType {
				temp, err := strconv.Atoi(src.(string))
				if err != nil {
					return nil, err
				}
				result := temp
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
