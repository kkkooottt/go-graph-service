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

type Instace struct {
	Chart []ChartPoint
}

func (i *Instace) PutValues(c []ChartPoint) {
	i.Chart = c
}

func New() *Instace {
	chart := make([]ChartPoint, 0)
	instance := Instace{
		Chart: chart,
	}
	return &instance
}

func (i *Instace) Start() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./index.html")
	})

	http.HandleFunc("/graph", func(w http.ResponseWriter, r *http.Request) {

		resp := ChartResponse{
			Points: i.Chart,
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
