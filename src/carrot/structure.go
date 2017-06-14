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

type Body struct {
	code			string
	fileType		string
	line			int
	column			int
	wordToComplete	string
	offset			int
}

type Data struct {
	body  Body
}

func GenMsg(line string) []byte {
	return []byte(`&carrot.Data{&carrot.Body{line, "python", 0, 1, line, len([]rune(line))+1}}`)
}

type Routine struct {
	SendTime    time.Time
	ReceiveTime time.Time
	Diff        time.Duration // milliseconds
	ReceivedMsg string
}
