package syslog

import (
	"github.com/jeromer/syslogparser"
)

type PassthruParser struct {
	buff []byte
}

func NewPassthruParser(buff []byte) *PassthruParser {
	return &PassthruParser{
		buff: buff,
	}
}

func (p *PassthruParser) Parse() error {
	return nil
}

func (p *PassthruParser) Dump() syslogparser.LogParts {
	return syslogparser.LogParts{
		"content":  string(p.buff),
	}
}
