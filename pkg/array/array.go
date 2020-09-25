package array

import (
	"reflect"
)

// InArray search an element in an array.
func InArray(value interface{}, array interface{}) (found bool, index int) {
	index = -1
	found = false

	switch reflect.Indirect(reflect.ValueOf(array)).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		sLen := s.Len()

		for i := 0; i < sLen; i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				index = i
				found = true

				break
			}
		}
	}
	return
}

// StringInSlice finds string in slice.
func StringInSlice(a string, slice []string) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}

// RemoveStringFromSlice removes string from slice.
func RemoveStringFromSlice(a string, slice []string) []string {
	for i, str := range slice {
		if str == a {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// StringArrayIntersection returns the intersection of two slices.
func StringArrayIntersection(arr1, arr2 []string) []string {
	arrMap := map[string]bool{}
	var result []string

	for _, value := range arr1 {
		arrMap[value] = true
	}

	for _, value := range arr2 {
		if arrMap[value] {
			result = append(result, value)
		}
	}

	return result
}

// RemoveDuplicatesFromStringArray removes duplicates strings from slice.
func RemoveDuplicatesFromStringArray(arr []string) []string {
	result := make([]string, 0, len(arr))
	seen := make(map[string]bool)

	for _, item := range arr {
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

// StringSliceDiff returns differences between two slices.
func StringSliceDiff(a, b []string) []string {
	m := make(map[string]bool)
	var result []string

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if !m[item] {
			result = append(result, item)
		}
	}
	return result
}
