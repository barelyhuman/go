package slices_test

import (
	"testing"

	"github.com/barelyhuman/go/slices"
)

func TestChunk(t *testing.T) {
	exampleSlice := []string{"hello", "all", "humans"}

	t.Run("single item chunks", func(t *testing.T) {
		chunkedSlice := slices.Chunk(exampleSlice, 1)
		if len(chunkedSlice) != 3 {
			t.Fail()
		}
	})

	t.Run("multiple item chunks", func(t *testing.T) {
		chunkedSlice := slices.Chunk(exampleSlice, 2)
		if len(chunkedSlice) != 2 {
			t.Fail()
		}
		if len(chunkedSlice[0]) != 2 {
			t.Fail()
		}
	})

}

func TestPickField(t *testing.T) {
	type pickFieldTestType struct {
		value string
	}

	exampleSlice := []pickFieldTestType{{value: "wake up"}, {value: "code"}, {value: "repeat"}}

	t.Run("with func", func(t *testing.T) {
		pickedSlices := slices.PickField(exampleSlice,
			slices.PickWithFunc(
				func(k pickFieldTestType) string {
					return k.value
				},
			))

		if len(pickedSlices) != len(exampleSlice) {
			t.Fail()
		}

		for i, sliceItem := range exampleSlice {
			if sliceItem.value != pickedSlices[i] {
				t.Fail()
			}
		}
	})

	t.Run("with key func", func(t *testing.T) {
		pickedSlices := slices.PickField(exampleSlice,
			slices.PickWithKey[pickFieldTestType]("value"),
		)

		if len(pickedSlices) != len(exampleSlice) {
			t.Fail()
		}

		for i, sliceItem := range exampleSlice {
			if sliceItem.value != pickedSlices[i].String() {
				t.Fail()
			}
		}
	})
}
