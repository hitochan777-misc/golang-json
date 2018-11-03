package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func makeType(v interface{}) reflect.Type {
	typ := reflect.TypeOf(v)
	var fields []reflect.StructField
	for i := 0; i < typ.NumField(); i++ {
		fields = append(fields, typ.Field(i))
	}
	return reflect.StructOf(fields)
}

func main() {
	typ := makeType(Person{})
	p := reflect.New(typ).Elem()
	p.FieldByName("Name").Set(reflect.ValueOf("hoge"))
	fmt.Println(p)
}
