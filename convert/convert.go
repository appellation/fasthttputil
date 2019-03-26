package convert

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrCannotConvert = errors.New("cannot convert given type")

func FromString(s string, t reflect.Type) (v reflect.Value, err error) {
	if len(s) == 0 {
		v = reflect.Zero(t)
		return
	}

	switch t.Kind() {
	case reflect.Bool:
		var b bool
		b, err = strconv.ParseBool(s)
		v = reflect.ValueOf(b)
	case reflect.Int:
		var i int64
		i, err = strconv.ParseInt(s, 0, 0)
		v = reflect.ValueOf(int(i))
	case reflect.Int8:
		var i int64
		i, err = strconv.ParseInt(s, 0, 8)
		v = reflect.ValueOf(int8(i))
	case reflect.Int16:
		var i int64
		i, err = strconv.ParseInt(s, 0, 16)
		v = reflect.ValueOf(int16(i))
	case reflect.Int32:
		var i int64
		i, err = strconv.ParseInt(s, 0, 32)
		v = reflect.ValueOf(int32(i))
	case reflect.Int64:
		var i int64
		i, err = strconv.ParseInt(s, 0, 64)
		v = reflect.ValueOf(i)
	case reflect.Uint:
		var i uint64
		i, err = strconv.ParseUint(s, 0, 0)
		v = reflect.ValueOf(uint(i))
	case reflect.Uint8:
		var i uint64
		i, err = strconv.ParseUint(s, 0, 8)
		v = reflect.ValueOf(uint8(i))
	case reflect.Uint16:
		var i uint64
		i, err = strconv.ParseUint(s, 0, 16)
		v = reflect.ValueOf(uint16(i))
	case reflect.Uint32:
		var i uint64
		i, err = strconv.ParseUint(s, 0, 32)
		v = reflect.ValueOf(uint32(i))
	case reflect.Uint64:
		var i uint64
		i, err = strconv.ParseUint(s, 0, 64)
		v = reflect.ValueOf(i)
	case reflect.Float32:
		var i float64
		i, err = strconv.ParseFloat(s, 32)
		v = reflect.ValueOf(float32(i))
	case reflect.Float64:
		var i float64
		i, err = strconv.ParseFloat(s, 64)
		v = reflect.ValueOf(i)
	case reflect.String:
		v = reflect.ValueOf(s)
	default:
		err = ErrCannotConvert
	}

	return
}
