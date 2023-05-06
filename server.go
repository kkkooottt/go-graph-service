package graph

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Reply struct {
	Type    string  `json:"type"`
	Data    Data    `json:"data"`
	Options Options `json:"options"`
}
type Datasets struct {
	Label       string    `json:"label"`
	Data        []float64 `json:"data"`
	BorderWidth int       `json:"borderWidth"`
}
type Data struct {
	Labels   []string   `json:"labels"`
	Datasets []Datasets `json:"datasets"`
}
type Y struct {
	BeginAtZero bool `json:"beginAtZero"`
}

type Scales struct {
	Y Y `json:"y"`
}
type Options struct {
	//Scales Scales `json:"scales"`
}

type Instace struct {
	Port  string
	Chart Reply
}

var (
	//go:embed static
	res embed.FS
)

const (
	TypeBar  = "bar"
	TypeLine = "line"
)

func New(port string) *Instace {

	instance := Instace{
		Port:  port,
		Chart: Reply{},
	}
	return &instance
}

func (i *Instace) PutValues(c Reply) {
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

		slb, err := json.Marshal(i.Chart)
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
