package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func readFile(file string) (dump string, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	return string(data), nil
}

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ErrorNil fails the test if an err is not nil.
func ErrorNil(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// ErrorNotNil fails the test if an err is not nil.
func ErrorNotNil(tb testing.TB, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: expected error but got none\n\n", filepath.Base(file), line)
		tb.FailNow()
	}
}

// Equals fails the test if expected is not equal to actual.
func Equals(tb testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, expected, actual)
		tb.FailNow()
	}
}
