package structs

import (
	"fmt"
	"reflect"
	"strconv"
)

// decoderFunc represents decoding functions for default built-in types.
type decoderFunc func(reflect.Value, interface{}) error

var (
	// List of built-in decoders indexed by their numeric constant values (eg: reflect.Bool = 1).
	decoders = []decoderFunc{
		reflect.Bool:          decodeBool,
		reflect.Int:           decodeInt,
		reflect.Int8:          decodeInt8,
		reflect.Int16:         decodeInt16,
		reflect.Int32:         decodeInt32,
		reflect.Int64:         decodeInt64,
		reflect.Uint:          decodeUint,
		reflect.Uint8:         decodeUint8,
		reflect.Uint16:        decodeUint16,
		reflect.Uint32:        decodeUint32,
		reflect.Uint64:        decodeUint64,
		reflect.Float32:       decodeFloat32,
		reflect.Float64:       decodeFloat64,
		reflect.Complex64:     decodeUnsupported,
		reflect.Complex128:    decodeUnsupported,
		reflect.Array:         decodeUnsupported,
		reflect.Chan:          decodeUnsupported,
		reflect.Func:          decodeUnsupported,
		reflect.Interface:     decodeUnsupported,
		reflect.Map:           decodeUnsupported,
		reflect.Ptr:           decodeUnsupported,
		reflect.Slice:         decodeSlice,
		reflect.String:        decodeString,
		reflect.Struct:        decodeUnsupported,
		reflect.UnsafePointer: decodeUnsupported,
	}

	// Global map of struct field specs that is populated once for every new
	// struct type that is scanned. This caches the field types and the corresponding
	// decoder functions to avoid iterating through struct fields on subsequent scans.
	globalStructMap = newStructMap()
)

func decodeBool(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case bool:
	    f.SetBool(s.(bool))
    case string:
        b, err := strconv.ParseBool(s.(string))
        if err != nil {
            return err
        }
        f.SetBool(b)
    default:
        return nil
    }
	return nil
}

func decodeInt8(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case int8,int16,int32,int64:
	    f.SetInt(s.(int64))
    case string:
        return decodeNumber(f, s.(string), 8)
    }
	return nil
}

func decodeInt16(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case int8,int16,int32,int64:
	    f.SetInt(s.(int64))
    case string:
        return decodeNumber(f, s.(string), 16)
    }
	return nil
}

func decodeInt32(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case int8,int16,int32,int64:
	    f.SetInt(s.(int64))
    case string:
        return decodeNumber(f, s.(string), 32)
    }
	return nil
}

func decodeInt64(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case int8,int16,int32,int64:
	    f.SetInt(s.(int64))
    case string:
        return decodeNumber(f, s.(string), 64)
    }
	return nil
}

func decodeInt(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case int8,int16,int32,int64:
	    f.SetInt(s.(int64))
    case string:
        return decodeNumber(f, s.(string), 0)
    }
	return nil
}

func decodeNumber(f reflect.Value, s string, bitSize int) error {
	v, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return err
	}
	f.SetInt(v)
	return nil
}

func decodeUint8(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case uint8,uint16,uint32,uint64:
	    f.SetUint(s.(uint64))
    case string:
        return decodeUnsignedNumber(f, s.(string), 8)
    }
	return nil
}

func decodeUint16(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case uint8,uint16,uint32,uint64:
	    f.SetUint(s.(uint64))
    case string:
        return decodeUnsignedNumber(f, s.(string), 16)
    }
	return nil
}

func decodeUint32(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case uint8,uint16,uint32,uint64:
	    f.SetUint(s.(uint64))
    case string:
        return decodeUnsignedNumber(f, s.(string), 32)
    }
	return nil
}

func decodeUint64(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case uint8,uint16,uint32,uint64:
	    f.SetUint(s.(uint64))
    case string:
        return decodeUnsignedNumber(f, s.(string), 64)
    }
	return nil
}

func decodeUint(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case uint8,uint16,uint32,uint64:
	    f.SetUint(s.(uint64))
    case string:
        return decodeUnsignedNumber(f, s.(string), 0)
    }
	return nil
}

func decodeUnsignedNumber(f reflect.Value, s string, bitSize int) error {
	v, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		return err
	}
	f.SetUint(v)
	return nil
}

func decodeFloat32(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case float32,float64:
	    f.SetFloat(s.(float64))
    case string:
        v, err := strconv.ParseFloat(s.(string), 32)
        if err != nil {
            return err
        }
        f.SetFloat(v)
    }
	return nil
}

// although the default is float64, but we better define it.
func decodeFloat64(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case float32,float64:
	    f.SetFloat(s.(float64))
    case string:
        v, err := strconv.ParseFloat(s.(string), 64)
        if err != nil {
            return err
        }
        f.SetFloat(v)
    }
	return nil
}

func decodeString(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case int8,int16,int32,int64,uint8,uint16,uint32,uint64:
	    f.SetString(fmt.Sprintf("%d", s))
    case float32,float64:
	    f.SetString(fmt.Sprintf("%f", s))
    case bool:
        if s.(bool) {
	        f.SetString("true")
        } else {
            f.SetString("false")
        }
    case string:
	    f.SetString(s.(string))
    }
	return nil
}

func decodeSlice(f reflect.Value, s interface{}) error {
    switch s.(type) {
    case []byte:
	    f.SetBytes(s.([]byte))
    case string:
        // []byte slice ([]uint8).
        if f.Type().Elem().Kind() == reflect.Uint8 {
            f.SetBytes([]byte(s.(string)))
        }
    }
	return nil
}

func decodeUnsupported(v reflect.Value, s interface{}) error {
	return fmt.Errorf("redis.Scan(unsupported %s)", v.Type())
}
