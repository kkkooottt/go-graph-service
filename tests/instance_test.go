package graph

import (
	"strconv"
	"testing"
)

func Test_Instance(t *testing.T) {

	grph := New()

	grph.Start()

	var chart []ChartPoint
	points := []float64{1, 3, 4.5, 5, 10}
	for k, v := range points {
		ch := ChartPoint{
			Label: strconv.Itoa(k),
			Value: v,
		}
		chart = append(chart, ch)
	}

	grph.PutValues(chart)

	c := make(chan bool)

	_ = <-c

}
