package main

import (
	"fmt"

	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/subscription"
)

type Wine struct {
	Name   string
	Age    int
	Rating float64
}

func GetWine() interface{} {
	w := []interface{}{
		Wine{"Merlot", 2011, 3.0},
		Wine{"Cabernet", 2010, 3.0},
		Wine{"Chardonnay", 2010, 4.0},
		Wine{"Pinot Grigio", 2009, 4.5},
	}

	return w
}

type Results map[int]Result

type Result struct {
	SumRating  float64
	NumSamples int
}

func Exec() (Results, <-chan subscription.Subscription) {

	results := make(Results)
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			wine, ok := item.(Wine)
			if ok {
				result := results[wine.Age]
				result.SumRating += wine.Rating
				result.NumSamples++
				results[wine.Age] = result
			}
		},
	}

	wine := GetWine()
	it, _ := iterable.New(wine)
	source := observable.From(it)
	sub := source.Subscribe(watcher)

	return results, sub
}

func main() {
	results, sub := Exec()
	<-sub
	for k, v := range results {
		fmt.Printf("Age: %d, Sample Size: %d, Average Rating: %.2f\n", k, v.NumSamples, v.SumRating/float64(v.NumSamples))
	}
}
