package log

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"text_changer/fileutil"
)

// configLog keeps path to log file
type configLog struct {
	Path string
	Name string
}

var cfg configLog

// NewConfig is a constructor for configLog struct
func NewConfig() *configLog {
	return &cfg
}

// level is a type for log importance hierarchy
type level int

// Importance hierarchy
const (
	FATAL level = iota
	WARNING
	INFO
)

// Default log messafe is INFO
var currentLevel = INFO

// String views of importance hierarchy
var levelWords = map[level]string{
	FATAL:   "[FATAL]",
	WARNING: "[WARNING]",
	INFO:    "[INFO]",
}

// SetLevel is a setter for log level
func SetLevel(l level) {
	currentLevel = l
}

// Print creates log message
func Print(l level, values ...interface{}) {
	if l > currentLevel {
		return
	}
	if l > INFO {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(os.Stdout)
	}

	// Creating message string
	buf := strings.TrimSuffix(fmt.Sprintln(append([]interface{}{levelWords[l] + ":"}, values...)...), "\n")

	// Adding message to log file
	isExist, err := fileutil.IsLogExist(string(os.PathSeparator)+cfg.Name, &cfg.Path)
	if err != nil {
		os.Stderr.WriteString(": " + "не удалось определить наличие файла логов")
	}
	if !isExist {
		err := fileutil.SaveLog([]byte(fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))+" "+buf), &cfg.Path, cfg.Name)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
	} else {
		err := fileutil.AddToLogFile(cfg.Path, cfg.Name, fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))+" "+buf)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
	}
	if l == FATAL {
		os.Exit(1)
	}
}

// PrintDuration prints the duration of some process
func PrintDuration(startTime time.Time) {
	Print(INFO, fmt.Sprintf("Finished in an %v", time.Since(startTime)))
}
