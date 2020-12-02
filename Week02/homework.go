package main

import (
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

var errMy = errors.New("sql.ErrNoRows")
var errCommont = errors.New("sql error")

func dao() error {
	err := mysqlServer()
	if err != nil {
		return err
	}
	return nil
}

func biz() error {
	err := dao()
	if err != nil {
		return xerrors.Wrap(err, "failed")
	}
	return nil
}

func service() error {
	return biz()
}

func mysqlServer() error {
	return errMy
}

func main() {
	err := service()
	fmt.Printf("service: %+v\n", err)
}

