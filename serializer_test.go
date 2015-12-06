package serializer

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	strSep = "~!@"
)

func TestEncodeIntArray(t *testing.T) {
	intArray := []int{1, 2, 3}
	if s, err := EncodeIntArray(intArray); err != nil {
		t.Error(err)
	} else if s != "1,2,3" {
		t.Error("Result wrong.")
	}

	if s, err := EncodeIntArray(emptyIntArray); err != nil {
		t.Error(err)
	} else if s != emptyString {
		t.Error("Result wrong.")
	}

	if _, err := EncodeIntArray(nil); err == nil {
		t.Error("Should pop an error.")
	}
}

func TestDecodeIntArray(t *testing.T) {
	s := "1,2,3"
	wrongS := "ab,1,1"

	intArray := []int{1, 2, 3}
	if result, err := DecodeIntArray(s); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(result, intArray) {
		t.Error("Result wrong.")
	}

	if result, err := DecodeIntArray(emptyString); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(result, emptyIntArray) {
		t.Error("Result wrong.")
	}

	if _, err := DecodeIntArray(wrongS); err == nil {
		t.Error("Should pop an error.")
	}
}

func TestEncodeStringArray(t *testing.T) {
	strArray := []string{"a", "b", "c"}
	result := fmt.Sprintf("a%sb%sc", strSep, strSep)
	if s, err := EncodeStringArray(strArray, strSep); err != nil {
		t.Error(err)
	} else if *s != result {
		t.Error("Result wrong.")
	}

	if s, err := EncodeStringArray(emptyStringArray, strSep); err != nil {
		t.Error(err)
	} else if *s != emptyString {
		t.Error("Result wrong.")
	}

	if _, err := EncodeStringArray(nil, strSep); err == nil {
		t.Error("Should pop an error.")
	}
}

func TestDecodeStringArray(t *testing.T) {
	strArray := []string{"a", "b", "c"}
	str := fmt.Sprintf("a%sb%sc", strSep, strSep)

	if result, err := DecodeStringArray(&str, strSep); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(result, strArray) {
		t.Error("Result wrong.")
	}

	empty := emptyString
	if result, err := DecodeStringArray(&empty, strSep); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(result, emptyStringArray) {
		t.Error("Result wrong.")
	}

	if _, err := DecodeStringArray(nil, strSep); err == nil {
		t.Error("Should pop an error.")
	}
}
