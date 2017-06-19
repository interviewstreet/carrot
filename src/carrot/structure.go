package carrot

import (
	"time"
)

type Base struct {
	URL, Proto string
	Count      int
	Msg        []byte
	Delay      int
	TickDelay  int
	Path       string
}

func GenMsg(line string) []byte {
	return []byte(`{"body":{"code":"`+line+`","fileType":"python","line":0,"column":`+string(len([]rune(line)))+`,"wordToComplete":"`+line+`","offset":`+string(len([]rune(line))+1)+ `}}`)
}

type Routine struct {
	SendTime    time.Time
	ReceiveTime time.Time
	Diff        time.Duration // milliseconds
	ReceivedMsg string
}
