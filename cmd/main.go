package main

import (
	"carrot"
	"fmt"
	"runtime"
	"time"
	"os"
	"bufio"
)

var msg = []byte(`{"body":{"code":"i","fileType":"python","line":0,"column":1,"wordToComplete":"i","offset":2}}`)
var count = 1000
var httpPort = 8900
var lines []string

func readLines() ([]string) {
  file, err := os.Open("input.txt")
  if err != nil {
    os.Exit(2)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if scanner.Err() != nil {
  	os.Exit(1)
  }
  return lines
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	latency := make(chan []float64)
	timeSeries := make(chan []time.Time)

	lines := readLines()

	currentTest := &carrot.Base{"localhost:8000", "ws", 100, msg, 10, 10, "/"}
	carrot.LoadTest(currentTest, latency, timeSeries, lines)

	data := <-latency
	timeData := <-timeSeries
	fmt.Println(data, timeData)
	fmt.Println("Running HTTP Server, Check /latency route at Port", httpPort)
	carrot.StartHTTPServer("8900", data, timeData)
	fmt.Scanln()
}
