package carrot

import (
	"time"
	"fmt"
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
	return []byte(fmt.Sprintf("%s%s%s%d%s%s%s%d%s", `{"body":{"code":"`, line, `","fileType":"python","line":0,"column":`, len([]rune(line)),`,"wordToComplete":"`, line, `","offset":`, len([]rune(line))+1, `}}`))
}

type Routine struct {
	SendTime    time.Time
	ReceiveTime time.Time
	Diff        time.Duration // milliseconds
	ReceivedMsg string
}
