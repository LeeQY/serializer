package serializer

import (
	"errors"
	"strconv"
	"strings"
)

const (
	emptyString = ""
	intSep      = ","
)

var (
	emptyIntArray    = make([]int, 0)
	emptyStringArray = make([]string, 0)
)

// Encode the int array to one string seperate by ","
func EncodeIntArray(array []int) (string, error) {
	if array == nil {
		return emptyString, errors.New("Array can't be nil.")
	}

	l := len(array)
	if l == 0 {
		return emptyString, nil
	}

	strArray := make([]string, l)
	for i := 0; i < l; i++ {
		strArray[i] = strconv.Itoa(array[i])
	}

	return strings.Join(strArray, intSep), nil
}

// Decode the string to int array split by ","
func DecodeIntArray(s string) ([]int, error) {
	if l := len(s); l == 0 {
		return emptyIntArray, nil
	}

	strArray := strings.Split(s, intSep)
	intArray := make([]int, len(strArray))

	var err error
	for i := 0; i < len(strArray); i++ {
		if intArray[i], err = strconv.Atoi(strArray[i]); err != nil {
			return nil, err
		}
	}
	return intArray, nil
}

// Encode the string array to one string with seperater.
// The seperater should be very rarely used string in the world like "&^%".
func EncodeStringArray(array []string, sep string) (*string, error) {
	if array == nil {
		return nil, errors.New("Array can't be nil.")
	}

	if len(sep) == 0 {
		return nil, errors.New("Seperater can't be empty to encode.")
	}

	if l := len(array); l == 0 {
		empty := emptyString
		return &empty, nil
	}

	result := strings.Join(array, sep)
	return &result, nil
}

// Decode the string to string array with seperater.
// The seperater should be the same with the one in Encode method.
func DecodeStringArray(s *string, sep string) ([]string, error) {
	if s == nil {
		return nil, errors.New("string can't be nil.")
	}

	if len(sep) == 0 {
		return nil, errors.New("Seperater can't be empty to encode.")
	}

	if *s == "" {
		return emptyStringArray, nil
	}

	return strings.Split(*s, sep), nil
}
