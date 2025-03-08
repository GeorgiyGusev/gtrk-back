package db_sp_call

import (
	"fmt"
)

type DBError struct {
	hasErr         bool
	noData         bool
	Code           string
	Message        string
	AdditionalData string
}

func (e DBError) NoData() bool {
	return e.noData
}

func (e DBError) HasErr() bool {
	return e.hasErr
}

func (e DBError) Error() string {
	return fmt.Sprintf("Err: %s, detail: %s", e.Code, e.Message)
}
