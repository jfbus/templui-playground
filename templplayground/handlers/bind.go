package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func Bind(r *http.Request, obj any) error {
	val := reflect.ValueOf(obj).Elem()
	valType := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanSet() {
			continue
		}
		name := valType.Field(i).Name
		value := r.FormValue(name)
		if value == "" {
			continue
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Bool:
			field.SetBool(value == "true")
		case reflect.Int:
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("%s: cannot convert %w", name, err)
			}
			field.SetInt(int64(intValue))
		}
	}
	return nil
}
