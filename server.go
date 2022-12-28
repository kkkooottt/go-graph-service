package graph

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ChartPoint struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}
type ChartResponse struct {
	Points []ChartPoint `json:"points"`
}

var Chart []ChartPoint

func PutValues(c []ChartPoint) {
	Chart = c
}

func Init() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./index.html")
	})

	http.HandleFunc("/graph", func(w http.ResponseWriter, r *http.Request) {
		/*
			var chart []ChartPoint
			points := []float64{1, 3, 4.5, 5, 10}
			for k, v := range points {
				ch := ChartPoint{
					Label: strconv.Itoa(k),
					Value: v,
				}
				chart = append(chart, ch)
			}

		*/
		resp := ChartResponse{
			Points: Chart,
		}
		slb, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
		}
		w.Write(slb)
	})

	fmt.Println("Server is listening...on 8080")
	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
	fmt.Println("Unexpected exit...")
}
