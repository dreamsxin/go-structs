package structs

import (
	"reflect"
	"strings"
	"sync"
)

// structMap contains the map of struct fields for target structs
// indexed by the struct type.
type structMap struct {
	m sync.Map
}

func newStructMap() *structMap {
	return new(structMap)
}

func (s *structMap) get(t reflect.Type) *structSpec {
	if v, ok := s.m.Load(t); ok {
		return v.(*structSpec)
	}

	spec := newStructSpec(t, "redis")
	s.m.Store(t, spec)
	return spec
}

//------------------------------------------------------------------------------

// structSpec contains the list of all fields in a target struct.
type structSpec struct {
	m map[string]*structField
}

func (s *structSpec) set(tag string, sf *structField) {
	s.m[tag] = sf
}

func newStructSpec(t reflect.Type, fieldTag string) *structSpec {
	numField := t.NumField()
	out := &structSpec{
		m: make(map[string]*structField, numField),
	}

	for i := 0; i < numField; i++ {
		f := t.Field(i)

		tag := f.Tag.Get(fieldTag)
		if tag == "-" {
			continue
		}
        if tag != "" {
            tag = strings.Split(tag, ",")[0]
            if tag == "" {
                continue
            }
        } else {
            tag = f.Name
        }

		// Use the built-in decoder.
		out.set(tag, &structField{index: i, fn: decoders[f.Type.Kind()]})
	}

	return out
}

//------------------------------------------------------------------------------

// structField represents a single field in a target struct.
type structField struct {
	index int
	fn    decoderFunc
}
