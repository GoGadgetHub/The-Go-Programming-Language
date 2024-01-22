package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	Name string
	Age  string
}

func main() {
	typ := reflect.Indirect(reflect.ValueOf(&Account{})).Type()
	fmt.Println(typ.Name()) // Account

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println(field.Name) // Username Password
	}
}
