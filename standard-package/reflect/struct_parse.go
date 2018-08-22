package main

import (
	"reflect"
	"log"
	"strings"
	"fmt"
)

type Point struct {
	X int `urlenc:"x,omitempty"`
	Y int `urlenc:"y,omitempty"`
}
func main() {
	res1 := parse(&Point{X: 10})
	fmt.Println(res1)

	res2 := parse(&Point{Y: 15})
	fmt.Println(res2)

	res3 := parse(&Point{X: 20, Y: 25})
	fmt.Println(res3)

}

func parse(p *Point) map[string]int64 {
	response := make(map[string]int64)
	t := reflect.TypeOf(*p)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}

		v, ok := f.Tag.Lookup("urlenc")
		if ok {
			// vの値がからであってもタグ指定が存在する
		} else {
			// タグ指定が存在しない
			continue
		}

		parts := strings.Split(v, ",")
		name := parts[0]
		omitempty := parts[1] // "omitempty" or ""

		rv := reflect.ValueOf(p).Elem()
		switch rv.Kind() {
		case reflect.Struct:
			fv := rv.Field(i)
			if omitempty == "omitempty" && fv.Int() == 0 {
				continue
			}
			response[name] = fv.Int()
		default:
			log.Println("not support")
		}
	}
	return response
}