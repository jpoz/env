package decoder

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	envTagKey    = "env"
	expandTagKey = "expand"
)

var (
	ErrNoStructPointer = errors.New("Must pass a struct pointer to Decode")
)

// Decode reads the current environment variables and stores variables values
// pointed to by v.
//
// Supported value types are: bool, int, int8, int16, int32, int64, string
// struct
//
func Decode(v interface{}) error {
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Ptr {
		return ErrNoStructPointer
	}

	typ = typ.Elem()
	if typ.Kind() != reflect.Struct {
		return ErrNoStructPointer
	}

	values := reflect.ValueOf(v).Elem()

	return traverse(typ, values)
}

func traverse(typ reflect.Type, values reflect.Value) error {
	err := NewError()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		val := values.FieldByName(field.Name)

		if val.IsValid() && val.CanSet() {
			var e error

			switch val.Kind() {
			case reflect.Bool:
				e = setBool(field, val)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				e = setInt(field, val)
			case reflect.String:
				e = setString(field, val)
			case reflect.Struct:
				e = traverse(field.Type, val)
			}

			if e != nil {
				err.Errors[field.Name] = e
			}
		}
	}

	if len(err.Errors) != 0 {
		return err
	}
	return nil
}

func setBool(field reflect.StructField, val reflect.Value) error {
	envVal := fieldValue(field)
	boolVal, err := convertToBool(envVal)
	if err == nil {
		val.SetBool(boolVal)
	}
	return err
}

func convertToBool(envVal *string) (boolVal bool, err error) {
	if envVal != nil {
		boolVal, err = strconv.ParseBool(*envVal)
	}
	return
}

func setString(field reflect.StructField, val reflect.Value) error {
	envVal := fieldValue(field)
	if envVal != nil {
		val.SetString(*envVal)
	}
	return nil
}

func setInt(field reflect.StructField, val reflect.Value) error {
	envVal := fieldValue(field)
	intVal, err := convertToInt64(envVal)
	if err == nil {
		val.SetInt(intVal)
	}
	return err
}

func convertToInt64(envVal *string) (intVal int64, err error) {
	if envVal != nil {
		intVal, err = strconv.ParseInt(*envVal, 0, 64)
	}
	return
}

func fieldValue(field reflect.StructField) *string {
	var envVal string

	envVal = field.Tag.Get(expandTagKey)
	if envVal != "" {
		envVal = os.ExpandEnv(envVal)
		goto CheckValue
	}

	envVal = field.Tag.Get(envTagKey)
	if envVal == "" {
		envVal = strings.ToUpper(field.Name)
	}
	envVal = os.Getenv(envVal)
	goto CheckValue

CheckValue:
	if envVal != "" {
		return &envVal
	}
	return nil
}
