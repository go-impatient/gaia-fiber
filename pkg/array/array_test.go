package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInArray tests InArray function.
func TestInArray(t *testing.T) {
	tInt := []int{10, 56, 23, 85}
	tString := []string{"45", "ghgh", "kl7878"}

	// Int
	// ---
	found, index := InArray(56, tInt)
	foundExpected, indexExpected := true, 1
	assert.True(t, found == foundExpected && index == indexExpected)

	found, index = InArray(589, tInt)
	foundExpected, indexExpected = false, -1
	assert.True(t, found == foundExpected && index == indexExpected)

	// String
	found, index = InArray("kl7878", tString)
	foundExpected, indexExpected = true, 2
	assert.True(t, found == foundExpected && index == indexExpected)

	found, index = InArray(589, tString)
	foundExpected, indexExpected = false, -1
	assert.True(t, found == foundExpected && index == indexExpected)
}

// TestStringArrayIntersection tests StringArrayIntersection function.
func TestStringArrayIntersection(t *testing.T) {
	a := []string{"abc", "def", "ghi"}
	b := []string{"jkl"}
	c := []string{"def"}

	assert.Empty(t, StringArrayIntersection(a, b))
	assert.Len(t, StringArrayIntersection(a, c), 1)
}

// TestRemoveDuplicatesFromStringArray tests RemoveDuplicatesFromStringArray function.
func TestRemoveDuplicatesFromStringArray(t *testing.T) {
	a := []string{"a", "b", "a", "a", "b", "c", "a"}

	assert.Len(t, RemoveDuplicatesFromStringArray(a), 3)
}

// TestStringSliceDiff tests StringSliceDiff function.
func TestStringSliceDiff(t *testing.T) {
	a := []string{"one", "two", "three", "four", "five", "six"}
	b := []string{"two", "seven", "four", "six"}
	expected := []string{"one", "three", "five"}

	assert.Equal(t, expected, StringSliceDiff(a, b))
}

// TestRemoveStringFromSlice tests RemoveStringFromSlice function.
func TestRemoveStringFromSlice(t *testing.T) {
	a := []string{"one", "two", "three", "four", "five", "six"}
	expected := []string{"one", "two", "three", "five", "six"}

	assert.Equal(t, expected, RemoveStringFromSlice("four", a))
}

// TestStringInSlice tests StringInSlice function.
func TestStringInSlice(t *testing.T) {
	a := []string{"one", "two", "three", "four", "five", "six"}

	assert.True(t, StringInSlice("four", a))
	assert.False(t, StringInSlice("seven", a))
}

func ExampleStringInSlice() {
	ok := StringInSlice("Coucou", []string{"Coucou", "Hello"})
	fmt.Println(ok)
	// Output:
	// true
}

func ExampleStringSliceDiff() {
	a := []string{"one", "two", "three", "four", "five", "six"}
	b := []string{"two", "seven", "four", "six"}
	diff := StringSliceDiff(a, b)

	fmt.Printf("diff: %v\n", diff)
	// Output:
	// diff: [one three five]
}
