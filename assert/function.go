package assert

import (
	"fmt"
	"reflect"
)

func formatMessage(msg []string) string {
	var output = ""
	switch len(msg) {
	case 0:
		output = ""
	case 1:
		output = msg[0]
	default:
		output = fmt.Sprintf(msg[0], msg[1:])
	}
	return output
}

func isBoolKind(kind reflect.Kind) bool {
	return kind == reflect.Bool
}

func isIntKind(kind reflect.Kind) bool {
	isInt := false
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		isInt = true
	}
	return isInt
}

func isUintKind(kind reflect.Kind) bool {
	isUint := false
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		isUint = true
	}
	return isUint
}

func isFloatKind(kind reflect.Kind) bool {
	isFloat := false
	switch kind {
	case reflect.Float32, reflect.Float64:
		isFloat = true
	}
	return isFloat
}

// IsEquals, tell the two params is equals or not
func IsEquals(a interface{}, b interface{}) bool {
	if a == nil || b == nil {
		return a == b
	}

	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)
	t1 := v1.Kind()
	t2 := v2.Kind()

	if isBoolKind(t1) && isBoolKind(t2) {
		return v1.Bool() == v2.Bool()
	}

	if isIntKind(t1) && isIntKind(t2) {
		return v1.Int() == v2.Int()
	}

	if isUintKind(t1) && isUintKind(t2) {
		return v1.Uint() == v2.Uint()
	}

	if isFloatKind(t1) && isFloatKind(t2) {
		return v1.Float() == v2.Float()
	}

	return reflect.DeepEqual(a, b)
}

// IsNil, tell the obj is nil or not
func IsNil(obj interface{}) bool {
	if obj == nil {
		return true
	}

	// nil must be a pointer, channel, func, interface, map, or slice type
	val := reflect.ValueOf(obj)
	return val.IsNil()
}