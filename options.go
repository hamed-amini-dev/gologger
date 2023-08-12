package gologger

import (
	"fmt"
	"os"
	"path/filepath"
)

type Option func(*myLogger) error

func OptionServiceName(str string) Option {
	return func(ml *myLogger) error {
		ml.config.serviceName = str
		return nil
	}
}

func OptionPrintToFile(filename string) Option {
	return func(ml *myLogger) error {
		ml.config.logFileName = filename

		dir := filepath.Dir(filename)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.Mkdir(dir, 0755)
			if err != nil {
				fmt.Printf("Not able to create directory. error: %v", err.Error())
				return err
			}
		}

		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Not able to open logfile for writing. error: %v", err.Error())
			return err
		}
		ml.l.SetOutput(f)

		return nil
	}

}

func OptionSetLevel(level Level) Option {
	return func(ml *myLogger) error {
		ml.config.level = level

		return nil
	}
}

func OptionSetFormatter(formatter Formatter) Option {
	return func(ml *myLogger) error {
		ml.l.SetFormatter(formatter)
		return nil
	}
}

// OptionReportCaller set the report caller, the offset or skip is depends on the development pattern.
// mostly set it to 3
func OptionReportCaller(offset int) Option {
	return func(ml *myLogger) error {
		ml.config.printCaller = offset
		// ml.l.SetReportCaller()
		return nil
	}
}
