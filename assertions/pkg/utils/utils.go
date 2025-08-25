package utils

import (
	"log"
	"reflect"
	"runtime"
)

var (
	// No Argument
	CheckError = DeferCheck
)

/*
*
Common Example:
connection.Close
*/
func DeferCheck(closeFunc func() error) {
	if err := closeFunc(); err != nil {
		log.Printf("Received Error; err, %v", err)
	}
}

func DeferCheckDebug(closeFunc func() error) {
	// Get the function's value
	funcValue := reflect.ValueOf(closeFunc)

	// Get the function's runtime.Func pointer
	runtimeFunc := runtime.FuncForPC(funcValue.Pointer())

	// Get the function's name
	if runtimeFunc != nil {
		log.Printf("Log Function name: %s", runtimeFunc.Name())
	} else {
		log.Println("Log Function name: <unknown>")
	}

	DeferCheck(closeFunc)
}

func CheckAppendError(appendFunc func(...any) error, row []string) {
	if err := appendFunc(row); err != nil {
		log.Printf("Received Error; err %v", err)
	}
}
