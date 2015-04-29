package env_struct

import (
	"fmt"
	"os"
	"reflect"
)

func Decode(obj interface{}) {
	objValues := reflect.ValueOf(obj).Elem()
	objType := reflect.TypeOf(obj).Elem()

	if objType.Kind() == reflect.Ptr {
		fmt.Printf("%v must be a pointer\n", obj)
		return
	}

	if objType.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", objType.Kind())
		return
	}

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		name := field.Tag.Get("env")

		objValue := objValues.FieldByName(field.Name)

		if name != "" {
			// If IsValid and CanSet are not true we will panic
			if objValue.IsValid() && objValue.CanSet() {

				// Set value
				in := objValue.Interface()
				switch in.(type) {
				case string:
					val := os.Getenv(name)
					if val == "" {
						val = field.Tag.Get("default")
					}
					objValue.SetString(val)
				}
			}
		}
	}
}
