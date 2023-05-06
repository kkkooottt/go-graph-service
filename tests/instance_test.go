package tests

import (
	"github.com/kkkooottt/go-graph-service"
	"testing"
)

func Test_Instance(t *testing.T) {

	t.Log("start")

	grph := graph.New("8080")

	t.Log("new instance inited")

	var chart graph.Reply

	chart.Type = graph.TypeLine // "bar" "line"

	app := graph.Datasets{
		Data:        []float64{1, 3, 4.5, 5, 10},
		Label:       "test",
		BorderWidth: 2,
	}

	app2 := graph.Datasets{
		Data:        []float64{10, 5, 3, 2, 1},
		Label:       "test2",
		BorderWidth: 2,
	}
	chart.Data.Datasets = append(chart.Data.Datasets, app)
	chart.Data.Datasets = append(chart.Data.Datasets, app2)
	chart.Data.Labels = []string{"1", "2", "3", "4", "5"}

	t.Log("test data prepared")

	grph.Start()

	t.Log("server started")
	grph.PutValues(chart)
	t.Log("data pushed")
	c := make(chan bool)

	_ = <-c

}
