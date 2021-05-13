package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
}

func createQuery(o order) string {
	s := fmt.Sprintf("insert into order values(%d, %d)", o.orderId, o.customerId)
	return s
}

func createQueryWithReflect(q interface{}) string {

	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)

	if reflect.TypeOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %d Type: %T, Value: %v\n", i, v.Field(i), v.Field(i))
		}
	}

	return ""
}

func main() {
	i := 10
	fmt.Printf("%d %T \n", i, i)

	o := order{
		orderId:    989384,
		customerId: 123,
	}

	fmt.Println(createQuery(o))

	e := employee{
		name:    "Havel",
		id:      873,
		address: "fall vally 24, istanborg Turkey",
		salary:  8700,
	}

	createQueryWithReflect(e)

}
