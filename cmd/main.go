package main

import (
	"carrot"
	"fmt"
	"runtime"
	"time"
	"sort"
)

var msg = []byte(`{"body":{"code":"i","fileType":"python","line":0,"column":1,"wordToComplete":"i","offset":2}}`)
var count = 1000
var httpPort = 8900

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	latency := make(chan []float64)
	timeSeries := make(chan []time.Time)

	lines := carrot.ReadLines()

	currentTest := &carrot.Base{"localhost:8000", "ws", 100, msg, 10, 10, "/"}
	carrot.LoadTest(currentTest, latency, timeSeries, lines)

	data := <-latency
	timeData := <-timeSeries
	fmt.Println(data, timeData)
	sort.Float64s(data)
	fmt.Println("99 percentile of the latency:", data[int(0.9*float64(len(data)))])
	fmt.Println("Running HTTP Server, Check /latency route at Port", httpPort)
	carrot.StartHTTPServer("8900", data, timeData)
	fmt.Scanln()
}
