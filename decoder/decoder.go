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

// Decode reads the current environment and stores it in the value pointed to
// by v.
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

	traverse(typ, values)

	return nil
}

func traverse(typ reflect.Type, values reflect.Value) {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		val := values.FieldByName(field.Name)

		if val.IsValid() && val.CanSet() {
			switch val.Kind() {
			case reflect.Bool:
				setBool(field, val)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				setInt(field, val)
			case reflect.String:
				setString(field, val)
			case reflect.Struct:
				traverse(field.Type, val)
			}
		}
	}
}

func setBool(field reflect.StructField, val reflect.Value) {
	envVal := fieldValue(field)
	boolVal, err := convertToBool(envVal)
	if err == nil {
		val.SetBool(boolVal)
	}
}

func convertToBool(envVal *string) (boolVal bool, err error) {
	if envVal != nil {
		boolVal, err = strconv.ParseBool(*envVal)
	}
	return
}

func setString(field reflect.StructField, val reflect.Value) {
	envVal := fieldValue(field)
	if envVal != nil {
		val.SetString(*envVal)
	}
}

func setInt(field reflect.StructField, val reflect.Value) {
	envVal := fieldValue(field)
	intVal, err := convertToInt64(envVal)
	if err == nil {
		val.SetInt(intVal)
	}
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
