package uhelpers

import (
	"reflect"
)

func StringKeysFromMap(m interface{}) []string {
	s := reflect.ValueOf(m)
	if s.Kind() != reflect.Map {
		panic("StringKeysFromMap() given a non-map type")
	}

	keys := make([]string, len(s.MapKeys()))
	i := 0
	for _, mapKey := range s.MapKeys() {
		key, ok := mapKey.Interface().(string)
		if !ok {
			panic("StringKeysFromMap() given a non-string key")
		}
		keys[i] = key
		i++
	}
	return keys
}
