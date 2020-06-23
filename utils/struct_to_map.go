package utils

import (
	"errors"
	"fmt"
	"reflect"
)

// 利用反射将结构体转化为map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func ArrayColumn(array interface{}, key string) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	t := reflect.TypeOf(array)
	v := reflect.ValueOf(array)
	if t.Kind() != reflect.Slice {
		return nil, errors.New("array type not slice")
	}
	if v.Len() == 0 {
		return nil, errors.New("array len is zero")
	}

	for i := 0; i < v.Len(); i++ {
		indexv := v.Index(i)
		if indexv.Type().Kind() != reflect.Struct {
			return nil, errors.New("element type not struct")
		}
		mapKeyInterface := indexv.FieldByName(key)
		if mapKeyInterface.Kind() == reflect.Invalid {
			return nil, errors.New("key not exist")
		}
		mapKeyString, err := interfaceToString(mapKeyInterface.Interface())
		if err != nil {
			return nil, err
		}
		result[mapKeyString] = indexv.Interface()
	}
	return result, err
}

func interfaceToString(v interface{}) (result string, err error) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		result = fmt.Sprintf("%v", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = fmt.Sprintf("%v", v)
	case reflect.String:
		result = v.(string)
	default:
		err = errors.New("can't transition to string")
	}
	return result, err
}
