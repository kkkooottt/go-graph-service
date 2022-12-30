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
	Port  string
	Chart []ChartPoint
}

func New(port string) *Instace {
	chart := make([]ChartPoint, 0)
	instance := Instace{
		Port:  port,
		Chart: chart,
	}
	return &instance
}

func (i *Instace) PutValues(c []ChartPoint) {
	i.Chart = c
}

func (i *Instace) Start() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./index.html")
	})

	http.HandleFunc("/graph", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("graph", i.Chart)
		resp := ChartResponse{
			Points: i.Chart,
		}
		slb, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
		}
		w.Write(slb)
	})
	go func() {
		fmt.Println("Server is listening...on ", i.Port)
		err := http.ListenAndServe(":"+i.Port, nil)

		fmt.Println(err)
		fmt.Println("Unexpected exit...")
	}()

}
