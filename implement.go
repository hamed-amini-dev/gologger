package gologger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type myLogger struct {
	l      *logrus.Logger
	config config
}

// New will cerate a new logger based on the options. it will not update the Default
func New(ops ...Option) (*myLogger, error) {
	lg := &myLogger{
		l:      logrus.New(),
		config: config{printCaller: -1},
	}
	for _, fn := range ops {
		err := fn(lg)
		if err != nil {
			return nil, err
		}
	}
	if lg.l.Out == nil {
		lg.l.SetOutput(os.Stdout)
	}
	if lg.config.level != 0 {
		lg.l.SetLevel(lg.config.level)
	}
	return lg, nil
}
