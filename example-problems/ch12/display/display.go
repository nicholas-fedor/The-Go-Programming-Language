// see page 333

// Package display provides a means to display structured data.
package display

import (
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T): \n", name, x)
	display(name, reflect.ValueOf(x))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		// Len method returns the number of elements of a slice or array value,
		// and Index(i) retrieves the element at index i, also as a reflect.Value;
		// it panics if i is out of bounds.
		for i := 0; i < v.Len(); i++ {
			// The display function recursively invokes itself on each element of the
			// sequence, appending the subscript notation "[i]" to the path.
			//
			// Although reflect.Value has many methods, only a few are safe to call
			// on any given value.
			// For example, the Index method may be called on values of kind Slice, Array,
			// or String, but panics for any other kind.
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		// The NumField method reports the number of fields in the struct, and Field(i) 
		// returns the value of the i-th field as a reflect.Value.
		// The list of fields includes ones promoted from anonymous fields.
		// To append the field selector notation ".f" to the path, we must obtain the 
		// reflect.Type of the struct and access the name of its i-th field.
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		// The MapKeys method returns a slice of reflect.Values, one per map key.
		// The order is undefined when iterating over a map.
		for _, key := range v.MapKeys() {
			// MapIndex(key) returns the value corresponding to key.
			// We append the subscript notation "[key]" to the path.
			// The type of a map key isn't restricted to the types formatAtom handles best;
			// arrays, structs, and interfaces can also be valid map keys.
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
			} else {
			// The Elem method returns the variable pointed to by a pointer as a reflect.Value.
			// This operation woudl be safe even if the pointer value is nil, in which case the 
			// result would have kind Invalid, but we use IsNil to detect nil pointers explicitly
			// so we can print a more appropriate message. 
			// We prefix the path with a "*" and parenthesize it to avoid ambiguity. 
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		// IsNil tests if the interface is nil.
		// If not nil, v.Elem() retrieves its dynamic value and we print its type and value.
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
			} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
