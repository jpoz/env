package env_struct

import (
	"errors"
	"os"
	"reflect"
	"strings"
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

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		// TODO check if struct

		envVar := field.Tag.Get("env")
		if envVar == "" {
			envVar = strings.ToUpper(field.Name)
		}
		expandString := field.Tag.Get("expand")

		val := values.FieldByName(field.Name)

		// If IsValid and CanSet are not true we will panic
		if val.IsValid() && val.CanSet() {
			if expandString != "" {
				in := val.Interface()
				switch in.(type) {
				case string:
					envVal := os.ExpandEnv(expandString)
					if envVal != "" {
						val.SetString(envVal)
					}
				}
			} else {
				in := val.Interface()
				switch in.(type) {
				case string:
					envVal := os.Getenv(envVar)
					if envVal != "" {
						val.SetString(envVal)
					}
				}
			}
		}
	}

	return nil
}
