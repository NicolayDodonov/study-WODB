package logger

import (
	"fmt"
	"io"
	"os"
	"study-WODB/internal/config"
	"time"
)

const (
	Debug = "[DEB] "
	Info  = "[INF] "
	Panic = "[PNC] "
	Error = "[ERR] "
	Fatal = "[FAT] "
)

const (
	DebugLevel LoggerType = 1
	InfoLevel  LoggerType = 2
	ErrorLevel LoggerType = 3
	OffLevel   LoggerType = 10
)

// LoggerType это уровень логгирования.
type LoggerType uint8

// Logger структура базового логгера, хранит в себе адрес хрангилища логгеров и уровень логгирования.
type Logger struct {
	file  *os.File
	level LoggerType
}

// New создаёт новый логгер и возвращает или указатель на него или ошибку.
func New(cnf *config.Config) (*Logger, error) {
	file, err := os.OpenFile(cnf.Logger.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return &Logger{}, err
	}
	return &Logger{
		file:  file,
		level: convert(cnf.Logger.Rang),
	}, nil
}

// =========- Методы логгирования -========= \\

func (logger *Logger) Debug(msg string) {
	if logger.level <= DebugLevel {
		t := time.Now()
		message := t.String() + " " + Debug + msg + "\n"
		fmt.Print(message)
		_, _ = io.WriteString(logger.file, message)

	}
}

func (logger *Logger) Info(msg string) {
	if logger.level <= InfoLevel {
		t := time.Now()
		message := t.String() + " " + Info + msg + "\n"
		fmt.Print(message)
		_, _ = io.WriteString(logger.file, message)
	}
}

func (logger *Logger) Panic(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+Panic+msg+"\n")
		panic(msg)
	}
}

func (logger *Logger) Error(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		message := t.String() + " " + Error + msg + "\n"
		fmt.Print(message)
		_, _ = io.WriteString(logger.file, message)
	}
}

func (logger *Logger) Fatal(msg string) {
	t := time.Now()
	_, _ = io.WriteString(logger.file, t.String()+" "+Fatal+msg+"\n")
	panic(msg)
}

// =========- Внутренние методы -========= \\

func convert(s string) LoggerType {
	switch s {
	case "Debug":
		return DebugLevel
	case "Info":
		return InfoLevel
	case "Error":
		return ErrorLevel
	case "Off":
		return OffLevel
	default:
		return OffLevel
	}
}
