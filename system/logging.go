package system

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
)

const (

	// PathBase is the location of all dominion src files
	PathBase = "/usr/local/dominion"
	// PathLogs is the location of all dominion log files
	PathLogs = "/usr/local/dominion/logs"
)

// log is where normal & debugging messages are dumped to
var logger *log.Logger
var closeFile func() error

// Setup initializes logging and signal handling
func Setup(logFilePath string) error {
	if logger != nil {
		return errors.New("system.Setup has already been called")
	}

	err := os.MkdirAll(path.Dir(logFilePath), 0755)
	if err != nil {
		return fmt.Errorf("failed to check ensure log directory \"%v\" exists: %w", path.Dir(logFilePath), err)
	}

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file \"%v\": %w", logFilePath, err)
	}
	closeFile := logFile.Close
	logger = log.New(logFile, "", log.LstdFlags)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := closeFile(); err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	return nil
}

func Panic(err error) {
	if closeFile != nil {
		if err := closeFile(); err != nil {
			panic(err)
		}
	}
	panic(err)
}

func Logf(fmt string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	logger.Printf(fmt, args...)
}

func LogRoutinef(routineName, fmts string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	prefix := fmt.Sprintf("Routine [%s]: ", routineName)
	logger.Printf(prefix+fmts, args...)
}

func LogRPCf(rpcName, fmts string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	prefix := fmt.Sprintf("RPC-%s	: ", rpcName)
	logger.Printf(prefix+fmts, args...)
}

func Errorf(fmts string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	err := fmt.Errorf(fmts, args...)
	logger.Printf("Error: %v", err)
}
