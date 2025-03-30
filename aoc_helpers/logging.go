package aoc_helpers

import (
	log "github.com/sirupsen/logrus"
)

func DebugLog(dbgFlag *bool, format string, args ...interface{}) {
	if *dbgFlag {
		log.Debugf(format, args...)
	}
}

func SetupLogging(dbgFlag, traceFlag *bool) {
	if *dbgFlag {
		log.SetLevel(log.DebugLevel)
	} else if *traceFlag {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
