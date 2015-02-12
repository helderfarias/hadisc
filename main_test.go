package main

import (
	"fmt"
	"testing"
)

func TestMapRange(t *testing.T) {

	m := make(map[string]string)
	m["oi"] = "olá"

	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}
}
