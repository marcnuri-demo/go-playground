package main

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"testing"
)

func TestDuration(t *testing.T) {
	test := strfmt.Duration(0)
	test.UnmarshalText([]byte("1hours"))
	fmt.Println(test.Value())

	test.UnmarshalText([]byte("1ns1us1ms1s1m1h1d1w"))
	fmt.Println(test.Value())

	fmt.Printf("IS valid? %v\n", strfmt.IsDuration("not a valid duration 1ns"))

	test = strfmt.Duration(10000000000)
	byteDuration, _ := test.MarshalJSON()
	fmt.Println(string(byteDuration))
	fmt.Println(test.Value())
}
