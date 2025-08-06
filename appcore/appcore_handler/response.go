package appcore_handler

import (
	"log"
	"reflect"
	"strings"
)

var pagination = [6]string{"Limit", "Page", "Sort", "TotalRows", "TotalPages", "Rows"}

type ResponseObject struct {
	Data any `json:"data"`
}

type ResponseCreated struct {
	Data struct {
		ID any `json:"id"`
	} `json:"data"`
}

type ResponseError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewResponseObject(obj any) ResponseObject {
	return ResponseObject{Data: obj}
}

func NewResponseObjectWithSensitiveData[T any](obj T, unSelectField []string, hideValueField []string) ResponseObject {
	data := fieldRule[T](&obj, unSelectField, hideValueField)
	return ResponseObject{Data: data}
}

func NewResponseCreated(id any) ResponseCreated {
	return ResponseCreated{
		Data: struct {
			ID any `json:"id"`
		}{
			ID: id,
		},
	}
}

// TODO: maybe delete this function
func NewResponseError(error, message string) ResponseError {
	return ResponseError{
		Error:   error,
		Message: message,
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
func fieldRule[T any](t *T, unSelectField []string, hideValueField []string) any {
	rt := reflect.TypeOf(*t)
	rv := reflect.ValueOf(*t)

	log.Println("rt.Kind()", rt.Kind())

	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		var sets []any
		for j := 0; j < rv.Len(); j++ {
			obj := rv.Index(j).Interface()

			rt1 := reflect.TypeOf(obj)
			rv1 := reflect.ValueOf(obj)

			set := make(map[string]interface{}, rt1.NumField()-len(unSelectField))
			for i := 0; i < rt1.NumField(); i++ {
				field := rt1.Field(i)
				jsonKey := field.Tag.Get("json")
				if !contains(unSelectField, jsonKey) {
					if contains(hideValueField, jsonKey) {
						text := rv1.Field(i).String()
						show := 3
						hide := len(text) - show
						set[jsonKey] = text[:show] + strings.Repeat("x", hide)
					} else {
						set[jsonKey] = rv1.Field(i).Interface()
					}
				}
			}
			//return set
			sets = append(sets, set)
		}
		return sets
	default:
		isPaginationStruct := true
		sets := make(map[string]interface{}, rt.NumField())
		for i := 0; i < rt.NumField(); i++ {
			if rt.Field(i).Name != pagination[i] {
				isPaginationStruct = false
				break
			}
			jsonKey := rt.Field(i).Tag.Get("json")
			jsonLabel := strings.Split(jsonKey, ",")[0]
			sets[jsonLabel] = rv.Field(i).Interface()
		}
		if isPaginationStruct {
			v := rv.Field(5).Interface()
			sets["rows"] = fieldRule(&v, unSelectField, hideValueField)
			return sets
		}

		set := make(map[string]interface{}, rt.NumField()-len(unSelectField))
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			jsonKey := field.Tag.Get("json")
			if !contains(unSelectField, jsonKey) {
				if contains(hideValueField, jsonKey) {
					text := rv.Field(i).String()
					show := 3
					hide := len(text) - show
					set[jsonKey] = text[:show] + strings.Repeat("x", hide)
				} else {
					set[jsonKey] = rv.Field(i).Interface()
				}
			}
		}
		return set
	}
}
