package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
)

func TestMicroTime(t *testing.T) {
	microToMarshal := v1.MicroTime{
		Time: time.Now(),
	}
	byteTime, _ := microToMarshal.MarshalJSON()
	fmt.Println(string(byteTime))

	unmarshalled := v1.MicroTime{}
	unmarshalled.UnmarshalJSON([]byte("\"2020-01-23T18:22:39.500131Z\""))
	fmt.Println(unmarshalled.Time.String())
}
