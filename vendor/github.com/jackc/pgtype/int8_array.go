// Code generated by erb. DO NOT EDIT.

package pgtype

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/jackc/pgio"
)

type Int8Array struct {
	Elements   []Int8
	Dimensions []ArrayDimension
	Status     Status
}

func (dst *Int8Array) Set(src interface{}) error {
	// untyped nil and typed nil interfaces are different
	if src == nil {
		*dst = Int8Array{Status: Null}
		return nil
	}

	if value, ok := src.(interface{ Get() interface{} }); ok {
		value2 := value.Get()
		if value2 != value {
			return dst.Set(value2)
		}
	}

	// Attempt to match to select common types:
	switch value := src.(type) {

	case []int16:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*int16:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []uint16:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*uint16:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []int32:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*int32:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []uint32:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*uint32:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []int64:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*int64:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []uint64:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*uint64:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []int:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*int:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []uint:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*uint:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			elements := make([]Int8, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = Int8Array{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []Int8:
		if value == nil {
			*dst = Int8Array{Status: Null}
		} else if len(value) == 0 {
			*dst = Int8Array{Status: Present}
		} else {
			*dst = Int8Array{
				Elements:   value,
				Dimensions: []ArrayDimension{{Length: int32(len(value)), LowerBound: 1}},
				Status:     Present,
			}
		}
	default:
		// Fallback to reflection if an optimised match was not found.
		// The reflection is necessary for arrays and multidimensional slices,
		// but it comes with a 20-50% performance penalty for large arrays/slices
		reflectedValue := reflect.ValueOf(src)
		if !reflectedValue.IsValid() || reflectedValue.IsZero() {
			*dst = Int8Array{Status: Null}
			return nil
		}

		dimensions, elementsLength, ok := findDimensionsFromValue(reflectedValue, nil, 0)
		if !ok {
			return fmt.Errorf("cannot find dimensions of %v for Int8Array", src)
		}
		if elementsLength == 0 {
			*dst = Int8Array{Status: Present}
			return nil
		}
		if len(dimensions) == 0 {
			if originalSrc, ok := underlyingSliceType(src); ok {
				return dst.Set(originalSrc)
			}
			return fmt.Errorf("cannot convert %v to Int8Array", src)
		}

		*dst = Int8Array{
			Elements:   make([]Int8, elementsLength),
			Dimensions: dimensions,
			Status:     Present,
		}
		elementCount, err := dst.setRecursive(reflectedValue, 0, 0)
		if err != nil {
			// Maybe the target was one dimension too far, try again:
			if len(dst.Dimensions) > 1 {
				dst.Dimensions = dst.Dimensions[:len(dst.Dimensions)-1]
				elementsLength = 0
				for _, dim := range dst.Dimensions {
					if elementsLength == 0 {
						elementsLength = int(dim.Length)
					} else {
						elementsLength *= int(dim.Length)
					}
				}
				dst.Elements = make([]Int8, elementsLength)
				elementCount, err = dst.setRecursive(reflectedValue, 0, 0)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if elementCount != len(dst.Elements) {
			return fmt.Errorf("cannot convert %v to Int8Array, expected %d dst.Elements, but got %d instead", src, len(dst.Elements), elementCount)
		}
	}

	return nil
}

func (dst *Int8Array) setRecursive(value reflect.Value, index, dimension int) (int, error) {
	switch value.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		if len(dst.Dimensions) == dimension {
			break
		}

		valueLen := value.Len()
		if int32(valueLen) != dst.Dimensions[dimension].Length {
			return 0, fmt.Errorf("multidimensional arrays must have array expressions with matching dimensions")
		}
		for i := 0; i < valueLen; i++ {
			var err error
			index, err = dst.setRecursive(value.Index(i), index, dimension+1)
			if err != nil {
				return 0, err
			}
		}

		return index, nil
	}
	if !value.CanInterface() {
		return 0, fmt.Errorf("cannot convert all values to Int8Array")
	}
	if err := dst.Elements[index].Set(value.Interface()); err != nil {
		return 0, fmt.Errorf("%v in Int8Array", err)
	}
	index++

	return index, nil
}

func (dst Int8Array) Get() interface{} {
	switch dst.Status {
	case Present:
		return dst
	case Null:
		return nil
	default:
		return dst.Status
	}
}

func (src *Int8Array) AssignTo(dst interface{}) error {
	switch src.Status {
	case Present:
		if len(src.Dimensions) <= 1 {
			// Attempt to match to select common types:
			switch v := dst.(type) {

			case *[]int16:
				*v = make([]int16, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*int16:
				*v = make([]*int16, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]uint16:
				*v = make([]uint16, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*uint16:
				*v = make([]*uint16, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]int32:
				*v = make([]int32, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*int32:
				*v = make([]*int32, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]uint32:
				*v = make([]uint32, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*uint32:
				*v = make([]*uint32, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]int64:
				*v = make([]int64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*int64:
				*v = make([]*int64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]uint64:
				*v = make([]uint64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*uint64:
				*v = make([]*uint64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]int:
				*v = make([]int, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*int:
				*v = make([]*int, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]uint:
				*v = make([]uint, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*uint:
				*v = make([]*uint, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			}
		}

		// Try to convert to something AssignTo can use directly.
		if nextDst, retry := GetAssignToDstType(dst); retry {
			return src.AssignTo(nextDst)
		}

		// Fallback to reflection if an optimised match was not found.
		// The reflection is necessary for arrays and multidimensional slices,
		// but it comes with a 20-50% performance penalty for large arrays/slices
		value := reflect.ValueOf(dst)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		switch value.Kind() {
		case reflect.Array, reflect.Slice:
		default:
			return fmt.Errorf("cannot assign %T to %T", src, dst)
		}

		if len(src.Elements) == 0 {
			if value.Kind() == reflect.Slice {
				value.Set(reflect.MakeSlice(value.Type(), 0, 0))
				return nil
			}
		}

		elementCount, err := src.assignToRecursive(value, 0, 0)
		if err != nil {
			return err
		}
		if elementCount != len(src.Elements) {
			return fmt.Errorf("cannot assign %v, needed to assign %d elements, but only assigned %d", dst, len(src.Elements), elementCount)
		}

		return nil
	case Null:
		return NullAssignTo(dst)
	}

	return fmt.Errorf("cannot decode %#v into %T", src, dst)
}

func (src *Int8Array) assignToRecursive(value reflect.Value, index, dimension int) (int, error) {
	switch kind := value.Kind(); kind {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		if len(src.Dimensions) == dimension {
			break
		}

		length := int(src.Dimensions[dimension].Length)
		if reflect.Array == kind {
			typ := value.Type()
			if typ.Len() != length {
				return 0, fmt.Errorf("expected size %d array, but %s has size %d array", length, typ, typ.Len())
			}
			value.Set(reflect.New(typ).Elem())
		} else {
			value.Set(reflect.MakeSlice(value.Type(), length, length))
		}

		var err error
		for i := 0; i < length; i++ {
			index, err = src.assignToRecursive(value.Index(i), index, dimension+1)
			if err != nil {
				return 0, err
			}
		}

		return index, nil
	}
	if len(src.Dimensions) != dimension {
		return 0, fmt.Errorf("incorrect dimensions, expected %d, found %d", len(src.Dimensions), dimension)
	}
	if !value.CanAddr() {
		return 0, fmt.Errorf("cannot assign all values from Int8Array")
	}
	addr := value.Addr()
	if !addr.CanInterface() {
		return 0, fmt.Errorf("cannot assign all values from Int8Array")
	}
	if err := src.Elements[index].AssignTo(addr.Interface()); err != nil {
		return 0, err
	}
	index++
	return index, nil
}

func (dst *Int8Array) DecodeText(ci *ConnInfo, src []byte) error {
	if src == nil {
		*dst = Int8Array{Status: Null}
		return nil
	}

	uta, err := ParseUntypedTextArray(string(src))
	if err != nil {
		return err
	}

	var elements []Int8

	if len(uta.Elements) > 0 {
		elements = make([]Int8, len(uta.Elements))

		for i, s := range uta.Elements {
			var elem Int8
			var elemSrc []byte
			if s != "NULL" || uta.Quoted[i] {
				elemSrc = []byte(s)
			}
			err = elem.DecodeText(ci, elemSrc)
			if err != nil {
				return err
			}

			elements[i] = elem
		}
	}

	*dst = Int8Array{Elements: elements, Dimensions: uta.Dimensions, Status: Present}

	return nil
}

func (dst *Int8Array) DecodeBinary(ci *ConnInfo, src []byte) error {
	if src == nil {
		*dst = Int8Array{Status: Null}
		return nil
	}

	var arrayHeader ArrayHeader
	rp, err := arrayHeader.DecodeBinary(ci, src)
	if err != nil {
		return err
	}

	if len(arrayHeader.Dimensions) == 0 {
		*dst = Int8Array{Dimensions: arrayHeader.Dimensions, Status: Present}
		return nil
	}

	elementCount := arrayHeader.Dimensions[0].Length
	for _, d := range arrayHeader.Dimensions[1:] {
		elementCount *= d.Length
	}

	elements := make([]Int8, elementCount)

	for i := range elements {
		elemLen := int(int32(binary.BigEndian.Uint32(src[rp:])))
		rp += 4
		var elemSrc []byte
		if elemLen >= 0 {
			elemSrc = src[rp : rp+elemLen]
			rp += elemLen
		}
		err = elements[i].DecodeBinary(ci, elemSrc)
		if err != nil {
			return err
		}
	}

	*dst = Int8Array{Elements: elements, Dimensions: arrayHeader.Dimensions, Status: Present}
	return nil
}

func (src Int8Array) EncodeText(ci *ConnInfo, buf []byte) ([]byte, error) {
	switch src.Status {
	case Null:
		return nil, nil
	case Undefined:
		return nil, errUndefined
	}

	if len(src.Dimensions) == 0 {
		return append(buf, '{', '}'), nil
	}

	buf = EncodeTextArrayDimensions(buf, src.Dimensions)

	// dimElemCounts is the multiples of elements that each array lies on. For
	// example, a single dimension array of length 4 would have a dimElemCounts of
	// [4]. A multi-dimensional array of lengths [3,5,2] would have a
	// dimElemCounts of [30,10,2]. This is used to simplify when to render a '{'
	// or '}'.
	dimElemCounts := make([]int, len(src.Dimensions))
	dimElemCounts[len(src.Dimensions)-1] = int(src.Dimensions[len(src.Dimensions)-1].Length)
	for i := len(src.Dimensions) - 2; i > -1; i-- {
		dimElemCounts[i] = int(src.Dimensions[i].Length) * dimElemCounts[i+1]
	}

	inElemBuf := make([]byte, 0, 32)
	for i, elem := range src.Elements {
		if i > 0 {
			buf = append(buf, ',')
		}

		for _, dec := range dimElemCounts {
			if i%dec == 0 {
				buf = append(buf, '{')
			}
		}

		elemBuf, err := elem.EncodeText(ci, inElemBuf)
		if err != nil {
			return nil, err
		}
		if elemBuf == nil {
			buf = append(buf, `NULL`...)
		} else {
			buf = append(buf, QuoteArrayElementIfNeeded(string(elemBuf))...)
		}

		for _, dec := range dimElemCounts {
			if (i+1)%dec == 0 {
				buf = append(buf, '}')
			}
		}
	}

	return buf, nil
}

func (src Int8Array) EncodeBinary(ci *ConnInfo, buf []byte) ([]byte, error) {
	switch src.Status {
	case Null:
		return nil, nil
	case Undefined:
		return nil, errUndefined
	}

	arrayHeader := ArrayHeader{
		Dimensions: src.Dimensions,
	}

	if dt, ok := ci.DataTypeForName("int8"); ok {
		arrayHeader.ElementOID = int32(dt.OID)
	} else {
		return nil, fmt.Errorf("unable to find oid for type name %v", "int8")
	}

	for i := range src.Elements {
		if src.Elements[i].Status == Null {
			arrayHeader.ContainsNull = true
			break
		}
	}

	buf = arrayHeader.EncodeBinary(ci, buf)

	for i := range src.Elements {
		sp := len(buf)
		buf = pgio.AppendInt32(buf, -1)

		elemBuf, err := src.Elements[i].EncodeBinary(ci, buf)
		if err != nil {
			return nil, err
		}
		if elemBuf != nil {
			buf = elemBuf
			pgio.SetInt32(buf[sp:], int32(len(buf[sp:])-4))
		}
	}

	return buf, nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int8Array) Scan(src interface{}) error {
	if src == nil {
		return dst.DecodeText(nil, nil)
	}

	switch src := src.(type) {
	case string:
		return dst.DecodeText(nil, []byte(src))
	case []byte:
		srcCopy := make([]byte, len(src))
		copy(srcCopy, src)
		return dst.DecodeText(nil, srcCopy)
	}

	return fmt.Errorf("cannot scan %T", src)
}

// Value implements the database/sql/driver Valuer interface.
func (src Int8Array) Value() (driver.Value, error) {
	buf, err := src.EncodeText(nil, nil)
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}

	return string(buf), nil
}