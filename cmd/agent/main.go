package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/404notfoundhard/http-metric.git/internal/myMetrics"
)

func preprocessMetrik(metr myMetrics.Metrics) ([]string, error) {
	values := reflect.ValueOf(metr)
	types := values.Type()
	var value string
	result := make([]string, 0)
	for i := 0; i < values.NumField(); i++ {
		if types.Field(i).Name == "PollCount" {
			value = fmt.Sprintf("%d", values.Field(i).Uint())
			url := "http://localhost:8080/update/counter" + "/" + types.Field(i).Name + "/" + value
			result = append(result, url)
		} else {
			switch {
			case values.Field(i).CanFloat():
				value = strconv.FormatFloat(values.Field(i).Float(), 'f', 6, 64)
			case values.Field(i).CanInt():
				value = fmt.Sprintf("%d", values.Field(i).Int())
			case values.Field(i).CanUint():
				value = strconv.FormatUint(values.Field(i).Uint(), 10)
			default:
				url := "http://localhost:8080/notfound"
				fmt.Println(url)
				log.Fatal(url)
				return nil, fmt.Errorf(url)
			}
			url := "http://localhost:8080/update/gauge" + "/" + types.Field(i).Name + "/" + value
			result = append(result, url)
		}
	}
	return result, nil
}

func SendMetrics(metr myMetrics.Metrics, f func(metr myMetrics.Metrics) ([]string, error)) {
	urls, err := f(metr)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range urls {
		resp, err := http.Post(v, "Content-Type: text/plain", bytes.NewBuffer([]byte{0}))
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != http.StatusOK {
			log.Fatal(resp.StatusCode)
		}
		log.Print(resp.StatusCode)
	}
}

func main() {
	var mycounter uint64
	for {
		mycounter++
		my_metrics := myMetrics.Metrics{}
		my_metrics = my_metrics.ReadMetrics()
		my_metrics.PollCount = mycounter
		SendMetrics(my_metrics, preprocessMetrik)
		time.Sleep(22 * time.Second)
	}
}
