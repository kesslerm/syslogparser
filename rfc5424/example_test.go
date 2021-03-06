package rfc5424_test

import (
	"testing"

	"github.com/bruceadowns/syslogparser/rfc5424"
)

func TestExampleNewParser(t *testing.T) {
	b := `<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] An application event log entry...`
	buff := []byte(b)

	p := rfc5424.NewParser(buff)
	if err := p.Parse(); err != nil {
		t.Fatal(err)
	}

	t.Log(p.Dump())
}
