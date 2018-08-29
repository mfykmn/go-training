package main

import (
	"errors"
	"log"
	"reflect"
)

type User struct {
	ID int
}

func main() {
	m := map[string]string{
		"0":   "hello",
		"aaa": "world",
		"bbb": "mafuyuk",
	}
	if _, err := marshall(m); err != nil {
		log.Fatal(err)
	}

	s := struct {
		string
	}{
		"aaaaaaa",
	}
	if _, err := marshall(s); err != nil {
		log.Fatal(err)
	}

	u := &User{ID: 9}
	if _, err := marshall(u); err != nil {
		log.Fatal(err)
	}
}

func marshall(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		// map用のコード
		log.Println("is map")
		for _, key := range rv.MapKeys() {
			mv := rv.MapIndex(key)

			log.Println(mv.String())
		}
		return nil, nil
	case reflect.Struct:
		// struct用のコード
		log.Println("is struct")
		rt := rv.Type()
		for i := 0; i < rt.NumField(); i++ {
			ftv := rt.Field(i) // こっちはFiled
			log.Println(ftv.Name)

			fv := rv.Field(i) // こっちはValue
			log.Println(fv.String())
		}

		return nil, nil
	default:
		return nil, errors.New("urlenc.Marshal: unsupported type (" + rv.Type().String() + ")")
	}
}
