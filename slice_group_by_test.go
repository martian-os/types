package types

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSlice_GroupBy(t *testing.T) {
	t.Run("group by even/odd numbers", func(t *testing.T) {
		s := Slice[int]{1, 2, 3, 4, 5, 6}
		groups := s.GroupBy(func(i int) interface{} {
			return i % 2
		})

		expected := map[interface{}]Slice[int]{
			0: {2, 4, 6},
			1: {1, 3, 5},
		}

		if !reflect.DeepEqual(groups, expected) {
			t.Errorf("Expected %v, got %v", expected, groups)
		}
	})

	t.Run("group strings by first letter", func(t *testing.T) {
		s := Slice[string]{"apple", "banana", "apricot", "blueberry", "cherry"}
		groups := s.GroupBy(func(str string) interface{} {
			return string(str[0])
		})

		expected := map[interface{}]Slice[string]{
			"a": {"apple", "apricot"},
			"b": {"banana", "blueberry"},
			"c": {"cherry"},
		}

		if !reflect.DeepEqual(groups, expected) {
			t.Errorf("Expected %v, got %v", expected, groups)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		s := Slice[int]{}
		groups := s.GroupBy(func(i int) interface{} {
			return i % 2
		})

		if len(groups) != 0 {
			t.Errorf("Expected empty map, got %v", groups)
		}
	})

	t.Run("all elements map to the same key", func(t *testing.T) {
		s := Slice[int]{2, 4, 6, 8}
		groups := s.GroupBy(func(i int) interface{} {
			return 0
		})

		expected := map[interface{}]Slice[int]{
			0: {2, 4, 6, 8},
		}

		if !reflect.DeepEqual(groups, expected) {
			t.Errorf("Expected %v, got %v", expected, groups)
		}
	})

	t.Run("all elements map to different keys", func(t *testing.T) {
		s := Slice[int]{1, 2, 3}
		groups := s.GroupBy(func(i int) interface{} {
			return i
		})

		expected := map[interface{}]Slice[int]{
			1: {1},
			2: {2},
			3: {3},
		}

		if !reflect.DeepEqual(groups, expected) {
			t.Errorf("Expected %v, got %v", expected, groups)
		}
	})

	t.Run("preserves order within groups", func(t *testing.T) {
		s := Slice[int]{1, 2, 3, 4, 5, 6, 7, 8}
		groups := s.GroupBy(func(i int) interface{} {
			return i % 3
		})

		expected := map[interface{}]Slice[int]{
			0: {3, 6},
			1: {1, 4, 7},
			2: {2, 5, 8},
		}

		if !reflect.DeepEqual(groups, expected) {
			t.Errorf("Expected %v, got %v", expected, groups)
		}
	})
}

func ExampleSlice_GroupBy() {
	s := Slice[string]{"apple", "banana", "apricot", "blueberry", "cherry"}
	groups := s.GroupBy(func(str string) interface{} {
		return string(str[0])
	})

	// To ensure stable output, we'll sort the keys
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k.(string))
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %v\n", key, groups[key])
	}

	// Output:
	// a: [apple apricot]
	// b: [banana blueberry]
	// c: [cherry]
}
