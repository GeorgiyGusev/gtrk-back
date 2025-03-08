package db_sp_call

import (
	"fmt"
	"reflect"
	"time"
)

func timeDecoderHook(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	// Проверяем, нужно ли преобразовать строку в time.Time
	if from.Kind() == reflect.String && to == reflect.TypeOf(time.Time{}) {
		// Преобразуем строку в time.Time
		t, err := time.Parse(time.RFC3339, data.(string))
		if err != nil {
			return nil, fmt.Errorf("failed to parse time: %v", err)
		}
		return t, nil
	}

	// Ес че я рил писал комменты, это не джепете!!!
	return data, nil
}
