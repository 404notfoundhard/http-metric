// / сервер вываливается exit status 1 если брать хендлеры отсюда при post запросе, непонятно
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	myMetrics "internal/myMetrics"

	"github.com/go-chi/chi"
)

func SetValueHandle(m *myMetrics.Metrics) http.HandlerFunc {
	var err error
	return func(respWr http.ResponseWriter, r *http.Request) {
		fmt.Println(chi.URLParam(r, "name"))
		switch {
		case chi.URLParam(r, "name") == "GCCPUFraction":
			m.GCCPUFraction, err = strconv.ParseFloat(chi.URLParam(r, "value"), 64)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(m)
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
			fmt.Println(m)
			log.Fatal(chi.URLParam(r, "value"))
			log.Fatal("Unknown metric")
		}
	}
}

func GetAllValuesHandle(m *myMetrics.Metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := json.MarshalIndent(m, " ", " ")
		if err != nil {
			fmt.Println("error marshal json: ", err)
		}
		io.WriteString(w, string(res))
	}
}

func GetValueHandle(m *myMetrics.Metrics) http.HandlerFunc {
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
