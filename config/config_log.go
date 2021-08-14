package config

import (
	"bytes"
	"log"
)

var (
	buf           bytes.Buffer
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(&buf, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(&buf, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(&buf, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}
