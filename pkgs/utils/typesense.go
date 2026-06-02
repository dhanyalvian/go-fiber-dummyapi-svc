//- pkgs/utils/typesense.go

package utils

import (
	"reflect"
	"strings"
	"time"

	"github.com/typesense/typesense-go/v4/typesense/api"
	"github.com/typesense/typesense-go/v4/typesense/api/pointer"
)

func DeriveTypesenseFields[T any]() []api.Field {
	var entity T
	t := reflect.TypeOf(entity)

	return deriveFields(t, "")
}

func deriveFields(t reflect.Type, prefix string) []api.Field {
	return deriveFieldsInner(t, prefix, false)
}

func deriveFieldsInner(t reflect.Type, prefix string, insideArray bool) []api.Field {
	var fields []api.Field

	for i := range t.NumField() {
		f := t.Field(i)

		if !f.IsExported() {
			continue
		}

		if f.Anonymous {
			fields = append(fields, deriveFieldsInner(f.Type, prefix, insideArray)...)
			continue
		}

		tsTag := f.Tag.Get("typesense")

		if tsTag == "skip" || tsTag == "-" {
			continue
		}

		jsonName := jsonFieldName(f)
		if jsonName == "" || jsonName == "-" {
			continue
		}

		fullName := prefix + jsonName

		if f.Type.Kind() == reflect.Struct && f.Type != reflect.TypeOf(time.Time{}) {
			field := api.Field{Name: fullName, Type: "object"}
			applyFieldOptions(&field, tsTag)
			fields = append(fields, field)
			fields = append(fields, deriveFieldsInner(f.Type, fullName+".", insideArray)...)
			continue
		}

		if f.Type.Kind() == reflect.Slice && f.Type.Elem().Kind() == reflect.Struct {
			field := api.Field{Name: fullName, Type: "object[]"}
			applyFieldOptions(&field, tsTag)
			fields = append(fields, field)
			fields = append(fields, deriveFieldsInner(f.Type.Elem(), fullName+".", true)...)
			continue
		}

		tsType := resolveTypesenseType(f.Type)
		if insideArray {
			tsType = tsType + "[]"
		}
		field := api.Field{Name: fullName, Type: tsType}
		applyFieldOptions(&field, tsTag, insideArray)
		fields = append(fields, field)
	}

	return fields
}

func applyFieldOptions(field *api.Field, tsTag string, insideArray ...bool) {
	opts := parseTypesenseTag(tsTag)
	if hasOpt(opts, "facet") {
		field.Facet = pointer.True()
	}
	isInsideArray := len(insideArray) > 0 && insideArray[0]
	if hasOpt(opts, "sort") && !isInsideArray {
		field.Sort = pointer.True()
	}
	if hasOpt(opts, "optional") || isInsideArray {
		field.Optional = pointer.True()
	}
	if hasOpt(opts, "index") || hasOpt(opts, "facet") || hasOpt(opts, "sort") {
		field.Index = pointer.True()
	} else {
		field.Index = pointer.False()
	}
	if hasOpt(opts, "infix") {
		field.Infix = pointer.True()
	}
	if v, ok := opts["locale"]; ok {
		field.Locale = pointer.String(v)
	}
}

func jsonFieldName(f reflect.StructField) string {
	tag := f.Tag.Get("json")
	if tag == "" {
		return f.Name
	}
	parts := strings.Split(tag, ",")
	return parts[0]
}

func resolveTypesenseType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return "int32"
	case reflect.Int64:
		return "int64"
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.Slice:
		if t.Elem().Kind() == reflect.String {
			return "string[]"
		}
		return "string[]"
	default:
		return "string"
	}
}

func parseTypesenseTag(tag string) map[string]string {
	opts := make(map[string]string)
	if tag == "" {
		return opts
	}
	parts := strings.Split(tag, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		kv := strings.SplitN(p, "=", 2)
		if len(kv) == 2 {
			opts[kv[0]] = kv[1]
		} else {
			opts[p] = "true"
		}
	}
	return opts
}

func hasOpt(opts map[string]string, key string) bool {
	v, ok := opts[key]
	return ok && v == "true"
}
