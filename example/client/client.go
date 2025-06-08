package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

func DoSomething() error {
	err := someRequest()
	if err != nil {
		fwErr := fwsample.NewError(http.StatusNotFound, "user not found")
		return fwErr.Wrap(err)
	}
	return nil
}

func someRequest() error {
	return fmt.Errorf("user not found")
}

func main() {
	err := DoSomething()
	if err != nil {
		log.Println("Error occurred:", err)

		var fwErr *fwsample.FwError
		if errors.As(err, &fwErr) {
			fmt.Println("Error Code: ", fwErr.Code)
			fmt.Println("Error Message: ", fwErr.Message)
		}
	}
}
