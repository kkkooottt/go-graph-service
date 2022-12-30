package graph

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
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

var (
	//go:embed static
	res embed.FS
)

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
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		//http.ServeFile(writer, request, "./index.html")
		tpl, err := template.ParseFS(res, "static/index.html")
		if err != nil {
			log.Printf("page %s not found in pages cache...", request.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"userAgent": request.UserAgent(),
		}
		if err := tpl.Execute(w, data); err != nil {
			return
		}
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

	http.FileServer(http.FS(res))

	go func() {
		log.Println("Server is listening...on ", i.Port)
		err := http.ListenAndServe(":"+i.Port, nil)
		panic(err)
	}()

}
