package gologger

type config struct {
	serviceName string
	logFileName string
	level       Level
	printCaller int
}
