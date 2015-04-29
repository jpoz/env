package decoder

import (
	"errors"
	"os"
	"reflect"
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
//
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

	traverse(typ, values)

	return nil
}

func traverse(typ reflect.Type, values reflect.Value) {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		val := values.FieldByName(field.Name)

		switch val.Kind() {
		case reflect.String:
			setString(field, val)
		case reflect.Struct:
			traverse(field.Type, val)
		}
	}
}

func setString(field reflect.StructField, val reflect.Value) {
	envVar := field.Tag.Get(envTagKey)
	if envVar == "" {
		envVar = strings.ToUpper(field.Name)
	}
	expandString := field.Tag.Get(expandTagKey)

	// If IsValid and CanSet are not true we will panic
	if val.IsValid() && val.CanSet() {
		var envVal string
		if expandString != "" {
			envVal = os.ExpandEnv(expandString)
		} else {
			envVal = os.Getenv(envVar)
		}
		if envVal != "" {
			val.SetString(envVal)
		}
	}
}
