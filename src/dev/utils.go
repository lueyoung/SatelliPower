package main

import (
	"fmt"
	"reflect"
)

func foreach_struct(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%v -- %v\n", t.Field(k).Tag, v.Field(k).Interface())
	}
}
