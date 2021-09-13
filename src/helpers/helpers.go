package helpers

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"gorm.io/gorm"
)

// ParseStringToInt is a helper function to convert string to integer
func ParseStringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return result
}

// SliceContains is a helper function to check if slice of string contains string
func SliceContains(s interface{}, val interface{}) (int, bool) {
	slice := ConvertSliceToInterface(s)
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// convertSliceToInterface takes a slice passed in as an interface{}
func ConvertSliceToInterface(s interface{}) (slice []interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		return nil
	}
	length := v.Len()
	slice = make([]interface{}, length)
	for i := 0; i < length; i++ {
		slice[i] = v.Index(i).Interface()
	}
	return slice
}

// GetOffset is a helper function to get sql offset value from page and limit args
func GetOffset(page int, limit int) int {
	offset := (page - 1) * limit
	if offset < 0 {
		return 0
	}
	return offset
}

// NilToEmptyMap is a helper function to convert nil value to {}
func NilToEmptyMap(d *interface{}) interface{} {
	data := *d
	if *d == nil {
		data = make(map[string]interface{})
	}
	return data
}

// ParseStatusResponse is a helper function to get http status from error interface
func ParseStatusResponse(err error, s uint16) uint16 {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}
	return s
}

// ParseStatusResponse is a helper function to check if slice of struct has field with value
// same as like _.find(arr, (a) => a.field === value).length
func SliceOfStructContainsFieldValue(slice interface{}, fieldName string, fieldValueToCheck interface{}) bool {
	rangeOnMe := reflect.ValueOf(slice)
	for i := 0; i < rangeOnMe.Len(); i++ {
		s := rangeOnMe.Index(i)
		f := reflect.Indirect(s).FieldByName(fieldName)
		if f.IsValid() {
			if f.Interface() == fieldValueToCheck {
				return true
			}
		}
	}
	return false
}
