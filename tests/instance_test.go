package tests

import (
	"graph"
	"strconv"
	"testing"
)

func Test_Instance(t *testing.T) {
	t.Log("start")
	grph := graph.New("8080")
	t.Log("new instance inited")
	var chart []graph.ChartPoint
	points := []float64{1, 3, 4.5, 5, 10}
	for k, v := range points {
		ch := graph.ChartPoint{
			Label: strconv.Itoa(k),
			Value: v,
		}
		chart = append(chart, ch)
	}
	t.Log("test data prepared")

	grph.Start()

	t.Log("server started")
	grph.PutValues(chart)
	t.Log("data pushed")
	c := make(chan bool)

	_ = <-c

}
