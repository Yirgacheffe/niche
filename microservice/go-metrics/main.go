package main

import (
	"net/http"
	"time"

	metrics "github.com/rcrowley/go-metrics"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	c := metrics.GetOrRegisterCounter("counterhandler.counter", nil)
	c.Inc(1)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func TimerHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now()
	t := metrics.GetOrRegisterTimer("timeerhandler.timer", nil)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))

	t.UpdateSince(curTime)
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	t := metrics.GetOrRegisterTimer("reporthandler.writemetrics", nil)
	t.Time(func() {
		metrics.WriteJSONOnce(metrics.DefaultRegistry, w)
	})
}

func main() {

	http.HandleFunc("/counter", CounterHandler)
	http.HandleFunc("/timer", TimerHandler)
	http.HandleFunc("/report", ReportHandler)

	if err := http.ListenAndServe(":9001", nil); err != nil {
		panic(err)
	}

}
