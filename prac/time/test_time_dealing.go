package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second * 2)
	t1 := time.Now()
	fmt.Println("Time diff:", t1.Sub(t))

	skFormat := t.Format("01 January 2006")
	fmt.Println(skFormat)

	loc, _ := time.LoadLocation("Europe/Paris") // ignore error
	londonTime := t.In(loc)
	fmt.Println("Paris:", londonTime)
}
