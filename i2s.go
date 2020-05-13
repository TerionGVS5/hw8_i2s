package main

import (
	"fmt"
	"reflect"
)

func i2s(data interface{}, out interface{}) error {
	outStructOrSlice := reflect.Indirect(reflect.ValueOf(out))
	dataMapOrSlice := reflect.ValueOf(data)
	if outStructOrSlice.Kind() == reflect.Struct {
		iter := dataMapOrSlice.MapRange()
		for iter.Next() {
			keyData := iter.Key()
			valueData := iter.Value()
			currStructField := outStructOrSlice.FieldByName(keyData.String())
			if !currStructField.IsValid() {
				continue
			}
			switch currStructField.Kind() {
			case reflect.String:
				if valueData.Elem().Kind() != reflect.String {
					return fmt.Errorf("error when value to String")
				}
				currStructField.SetString(valueData.Elem().String())
			case reflect.Int:
				if valueData.Elem().Kind() != reflect.Float64 {
					return fmt.Errorf("error when value to Int")
				}
				currStructField.SetInt(int64(valueData.Elem().Float()))
			case reflect.Bool:
				if valueData.Elem().Kind() != reflect.Bool {
					return fmt.Errorf("error when value to Bool")
				}
				currStructField.SetBool(valueData.Elem().Bool())
			case reflect.Float64:
				if valueData.Elem().Kind() != reflect.Float64 {
					return fmt.Errorf("error when value to Float64")
				}
				currStructField.SetFloat(valueData.Elem().Float())
			}
		}
	}

	return nil
}
