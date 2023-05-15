package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/404notfoundhard/http-metric.git/internal/myMetrics"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func getValueHandle(m *myMetrics.Metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch {
		case chi.URLParam(r, "name") == "GCCPUFraction":
			fmt.Fprint(w, m.GCCPUFraction)
		case chi.URLParam(r, "name") == "Alloc":
			fmt.Fprint(w, m.Alloc)
		case chi.URLParam(r, "name") == "BuckHashSys":
			fmt.Fprint(w, m.BuckHashSys)
		case chi.URLParam(r, "name") == "Frees":
			fmt.Fprint(w, m.Frees)
		case chi.URLParam(r, "name") == "GCSys":
			fmt.Fprint(w, m.GCSys)
		case chi.URLParam(r, "name") == "HeapAlloc":
			fmt.Fprint(w, m.HeapAlloc)
		case chi.URLParam(r, "name") == "HeapIdle":
			fmt.Fprint(w, m.HeapIdle)
		case chi.URLParam(r, "name") == "HeapInuse":
			fmt.Fprint(w, m.HeapInuse)
		case chi.URLParam(r, "name") == "HeapObjects":
			fmt.Fprint(w, m.HeapObjects)
		case chi.URLParam(r, "name") == "HeapReleased":
			fmt.Fprint(w, m.HeapReleased)
		case chi.URLParam(r, "name") == "HeapSys":
			fmt.Fprint(w, m.HeapSys)
		case chi.URLParam(r, "name") == "LastGC":
			fmt.Fprint(w, m.LastGC)
		case chi.URLParam(r, "name") == "Lookups":
			fmt.Fprint(w, m.Lookups)
		case chi.URLParam(r, "name") == "MCacheInuse":
			fmt.Fprint(w, m.MCacheInuse)
		case chi.URLParam(r, "name") == "MCacheSys":
			fmt.Fprint(w, m.MCacheSys)
		case chi.URLParam(r, "name") == "MSpanInuse":
			fmt.Fprint(w, m.MSpanInuse)
		case chi.URLParam(r, "name") == "MSpanSys":
			fmt.Fprint(w, m.MSpanSys)
		case chi.URLParam(r, "name") == "Mallocs":
			fmt.Fprint(w, m.Mallocs)
		case chi.URLParam(r, "name") == "NextGC":
			fmt.Fprint(w, m.NextGC)
		case chi.URLParam(r, "name") == "NumForcedGC":
			fmt.Fprint(w, m.NumForcedGC)
		case chi.URLParam(r, "name") == "NumGC":
			fmt.Fprint(w, m.NumGC)
		case chi.URLParam(r, "name") == "OtherSys":
			fmt.Fprint(w, m.OtherSys)
		case chi.URLParam(r, "name") == "PauseTotalNs":
			fmt.Fprint(w, m.PauseTotalNs)
		case chi.URLParam(r, "name") == "StackInuse":
			fmt.Fprint(w, m.StackInuse)
		case chi.URLParam(r, "name") == "StackSys":
			fmt.Fprint(w, m.StackSys)
		case chi.URLParam(r, "name") == "Sys":
			fmt.Fprint(w, m.Sys)
		case chi.URLParam(r, "name") == "TotalAlloc":
			fmt.Fprint(w, m.TotalAlloc)
		case chi.URLParam(r, "name") == "RandomValue":
			fmt.Fprint(w, m.RandomValue)
		case chi.URLParam(r, "name") == "PollCount":
			fmt.Fprint(w, m.PollCount)
		default:
			http.Error(w, "Not found", http.StatusNotFound)
		}
	}
}

func getAllValuesHandle(m *myMetrics.Metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := json.MarshalIndent(m, " ", " ")
		if err != nil {
			fmt.Println("error marshal json: ", err)
		}
		io.WriteString(w, string(res))
	}
}

func setValueHandle(m *myMetrics.Metrics) http.HandlerFunc {
	return func(respWr http.ResponseWriter, r *http.Request) {
		switch {
		case chi.URLParam(r, "name") == "GCCPUFraction":
			m.GCCPUFraction, _ = strconv.ParseFloat(chi.URLParam(r, "value"), 64)
		case chi.URLParam(r, "name") == "Alloc":
			m.Alloc, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "BuckHashSys":
			m.BuckHashSys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "Frees":
			m.Frees, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "GCSys":
			m.GCSys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "HeapAlloc":
			m.HeapAlloc, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "HeapIdle":
			m.HeapIdle, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "HeapInuse":
			m.HeapInuse, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "HeapObjects":
			m.HeapObjects, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "HeapReleased":
			m.HeapReleased, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "HeapSys":
			m.HeapSys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "LastGC":
			m.LastGC, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "Lookups":
			m.Lookups, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "MCacheInuse":
			m.MCacheInuse, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "MCacheSys":
			m.MCacheSys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "MSpanInuse":
			m.MSpanInuse, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "MSpanSys":
			m.MSpanSys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "Mallocs":
			m.Mallocs, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "NextGC":
			m.NextGC, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "NumForcedGC":
			tmp, _ := strconv.ParseUint((chi.URLParam(r, "value")), 10, 32)
			m.NumForcedGC = uint32(tmp)
		case chi.URLParam(r, "name") == "NumGC":
			tmp, _ := strconv.ParseUint((chi.URLParam(r, "value")), 10, 32)
			m.NumGC = uint32(tmp)
		case chi.URLParam(r, "name") == "OtherSys":
			m.OtherSys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "PauseTotalNs":
			m.PauseTotalNs, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "StackInuse":
			m.StackInuse, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "StackSys":
			m.StackSys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "Sys":
			m.Sys, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "TotalAlloc":
			m.TotalAlloc, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		case chi.URLParam(r, "name") == "RandomValue":
			tmp, _ := strconv.ParseUint((chi.URLParam(r, "value")), 10, 32)
			m.RandomValue = uint32(tmp)
		case chi.URLParam(r, "name") == "PollCount":
			m.PollCount, _ = strconv.ParseUint((chi.URLParam(r, "value")), 10, 64)
		default:
			log.Fatal("Unknown metric")
		}
	}
}

func main() {
	r := chi.NewRouter()
	my_metrics := new(myMetrics.Metrics)
	// r.Use(middleware.requestID?)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)

	// myMetrika := new(myMetrics.Metrics)
	r.Get("/", getAllValuesHandle(my_metrics))
	r.Get("/value/{type}/{name}", getValueHandle(my_metrics))
	r.Post("/update/{type}/{name}/{value}", setValueHandle(my_metrics))
	log.Fatal(http.ListenAndServe(":8080", r))

}
