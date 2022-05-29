package slices

import (
	"log"
	"reflect"
)

// Chunk - create batches of an arbitrary typed slice into the given batch size
func Chunk[K comparable](slice []K, batchSize int) [][]K {
	var batches [][]K
	for i := 0; i < len(slice); i += batchSize {
		end := i + batchSize

		if end > len(slice) {
			end = len(slice)
		}

		batches = append(batches, slice[i:end])
	}

	return batches
}

type Picker[K interface{}, V comparable] func([]K) []V

// PickField - pick the value of a particular field from a slice into it's own slice using the given
// picker function
func PickField[K interface{}, V comparable](iterator []K, picker Picker[K, V]) []V {
	return picker(iterator)
}

// PickWithFunc
// get the key using a picker function
//
//  exampleSlice := []pickFieldTestType{{value: "wake up"}, {value: "code"}, {value: "repeat"}}
//
//  pickedSlices := slices.PickField(exampleSlice,
//  	slices.PickWithFunc(
//  		func(k pickFieldTestType) string {
//  			return k.value
//  		},
//  	))
func PickWithFunc[K interface{}, V comparable](pickerFunc func(K) V) Picker[K, V] {
	return func(k []K) (result []V) {
		for _, x := range k {
			result = append(result, pickerFunc(x))
		}
		log.Println(result)
		return result
	}
}

// PickWithKey
// when working with a single field, you can use this picker to get the values from the passed generic type.
//   exampleSlice := []pickFieldTestType{{value: "wake up"}, {value: "code"}, {value: "repeat"}}
//
//   pickedSlices := slices.PickField(exampleSlice,
//      slices.PickWithKey[pickFieldTestType]("value"),
//   )
func PickWithKey[K interface{}](field string) Picker[K, reflect.Value] {
	pickerFunc := func(k K) reflect.Value {
		r := reflect.ValueOf(k)
		f := reflect.Indirect(r).FieldByName(field)
		return f
	}
	return func(k []K) (result []reflect.Value) {
		for _, x := range k {
			result = append(result, pickerFunc(x))
		}
		return result
	}
}
