package logger

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

const (
	MessageIDKey = "MessageID"
	messageIDLen = 20
)

var (
	msgIDChars = []rune(`abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
)

func MessageID(msgID string) string {
	if len(msgID) > 0 {
		return msgID
	}
	return generateMsgID()
}

func generateMsgID() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	code := ""

	for i := 0; i < messageIDLen; i++ {
		charsIndex := generateRandom(0, len(msgIDChars)-1)
		code = fmt.Sprintf("%s%s", code, string(msgIDChars[charsIndex]))
	}

	return code
}

func generateRandom(min, max int) int {
	return min + rand.Intn(max-min)
}

const (
	logLevels = iota * 2
	infoLevel
	debugLevel
	diebugLevel
	warningLevel
	errorLevel
)

var logLevelsMap = map[int]string{
	infoLevel:    "INFO",
	debugLevel:   "DEBUG",
	diebugLevel:  "DIEBUG",
	warningLevel: "WARNING",
	errorLevel:   "ERROR",
}

func Info(msgID, message string) {
	generateLog(infoLevel, msgID, message, "", nil)
}

func Infof(msgID, message string, separator string, data ...interface{}) {
	generateLog(infoLevel, msgID, message, separator, data...)
}

func Debug(msgID, message string) {
	generateLog(debugLevel, msgID, message, "", nil)
}

func Debugf(msgID, message string, separator string, data ...interface{}) {
	generateLog(debugLevel, msgID, message, separator, data...)
}

func Diebug(msgID, message string) {
	generateLog(diebugLevel, msgID, message, "", nil)
}

func Diebugf(msgID, message string, separator string, data ...interface{}) {
	generateLog(diebugLevel, msgID, message, separator, data...)
}

func Warning(msgID, message string) {
	generateLog(warningLevel, msgID, message, "", nil)
}

func Warningf(msgID, message string, separator string, data ...interface{}) {
	generateLog(warningLevel, msgID, message, separator, data...)
}

func Error(msgID, message string) {
	generateLog(errorLevel, msgID, message, "", nil)
}

func Errorf(msgID, message string, separator string, data ...interface{}) {
	generateLog(errorLevel, msgID, message, separator, data...)
}

func generateLog(level int, msgID, message string, separator string, data ...interface{}) {
	var (
		msgStrWithData string
		logPrint       = fmt.Sprintf(`\e%s[%s]\e[0m msgID[ %s ] - %s`, colorByLabel(level), logLevelsMap[level], msgID, message)
	)

	for index, v := range data {
		if v != nil {

			if index == 0 {
				msgStrWithData = fmt.Sprintf("%s%v", msgStrWithData, v)
			} else {
				msgStrWithData = fmt.Sprintf("%s%s%v", msgStrWithData, separator, v)
			}
			if index == len(data)-1 {
				msgStrWithData = msgStrWithData + separator
			}
		}
	}

	logPrint = logPrint + msgStrWithData

	out, err := exec.Command("echo", "-e", logPrint).Output()

	if err != nil {
		Error("", err.Error())
	}

	if level == diebugLevel {
		log.Fatal(string(out))
	}

	log.Print(string(out))
}

func colorByLabel(level int) string {
	switch level {
	case infoLevel:
		return "[32m"
	case debugLevel:
		return "[36m"
	case diebugLevel:
		return "[36m"
	case warningLevel:
		return "[33m"
	case errorLevel:
		return "[31m"
	default:
		return "[37m"
	}
}
