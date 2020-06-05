package utils

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

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