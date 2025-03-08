package db_sp_call

import (
	"fmt"
	"github.com/GeorgiyGusev/gtrk-back/pkg/postgres"
	"github.com/goccy/go-json"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
	"log/slog"
	"reflect"
	"strings"
)

type DBCall struct {
	engine *sqlx.DB
}

func NewDBCall(dbConn *postgres.PostressConn) *DBCall {
	return &DBCall{engine: dbConn.DB}
}

func (d *DBCall) CallProcedure(scheme, name string, args interface{}) error {
	query, values, err := d.buildQuery(procedureCall, scheme, name, args)
	if err != nil {
		return err
	}

	_, err = d.engine.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

type dbResult struct {
	Data *string `db:"data"`
}

func (d *DBCall) CallFunction(dest interface{}, scheme, name string, args interface{}) (dberr DBError, err error) {
	query, values, err := d.buildQuery(functionCall, scheme, name, args)
	if err != nil {
		return dberr, err
	}

	var result dbResult
	if err = d.engine.Get(&result, query, values...); err != nil {
		return dberr, err
	}

	if result.Data == nil {
		return DBError{noData: true}, nil
	}

	var data map[string]interface{}
	if err = json.Unmarshal([]byte(*result.Data), &data); err != nil {
		return dberr, err
	}

	_, ok := data[keyErrorCode]
	if ok {
		return DBError{
			hasErr:  true,
			Code:    data[keyErrorCode].(string),
			Message: data[keyErrorMessage].(string),
		}, nil
	}

	return dberr, d.decode(dest, data)
}

func (d *DBCall) buildQuery(callWord, scheme, name string, args interface{}) (string, []any, error) {
	params, values, err := d.buildParams(args)
	if err != nil {
		return "", nil, err
	}
	slog.Info("buildQuery", "query", fmt.Sprintf("%s %s.%s(%s) AS data", callWord, scheme, name, params))
	return fmt.Sprintf("%s %s.%s(%s) AS data", callWord, scheme, name, params), values, nil
}

func (d *DBCall) buildParams(args interface{}) (string, []any, error) {
	if args == nil {
		return "", nil, nil
	}

	v := reflect.ValueOf(args)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("args must be a struct or pointer to a struct")
	}

	t := v.Type()

	params := make([]string, 0, t.NumField())
	values := make([]any, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get(tag)

		if tagValue == "" {
			continue
		}

		fieldValue := v.Field(i)
		params = append(params, fmt.Sprintf("_%s=>$%d", tagValue, len(params)+1))
		values = append(values, fieldValue.Interface())
	}

	return strings.Join(params, ","), values, nil
}

func (d *DBCall) decode(dest interface{}, res map[string]interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook:           mapstructure.ComposeDecodeHookFunc(timeDecoderHook),
		TagName:              tag,
		IgnoreUntaggedFields: true,
		Result:               dest,
	})
	if err != nil {
		return err
	}
	if dest == nil {
		return nil
	}
	return decoder.Decode(res)
}
