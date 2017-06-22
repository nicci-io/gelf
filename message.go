package gelf

import (
	"time"
)

const (
	GELF_VERSION = "1.1"
)

type Level byte

const (
	LEVEL_EMERGENCY Level = iota // 0
	LEVEL_ALERT                  // 1
	LEVEL_CRITICAL               // 2
	LEVEL_ERROR                  // 3
	LEVEL_WARNING                // 4
	LEVEL_NOTICE                 // 5
	LEVEL_INFO                   // 6
	LEVEL_DEBUG                  // 7
)


type Message struct {
	Version      string  `json:"version"`
	Host         string  `json:"host"`
	ShortMessage string  `json:"short_message"`
	Timestamp    float64 `json:"timestamp"`
	Level        Level   `json:"level"`
	FullMessage  string  `json:"full_message,omitempty"`
	Facility     string  `json:"facility,omitempty"`
	Line         uint    `json:"line,omitempty"`
	File         string  `json:"file,omitempty"`
}


func NewMessage(source string, level Level, shortMessage string) *Message {
	t := time.Now()
	timestamp := float64(t.Unix()) + float64(time.Second) / float64(t.Nanosecond())
	message := Message{
		Version:      GELF_VERSION,
		Host:         source,
		ShortMessage: shortMessage,
		Timestamp:    timestamp,
		Level:        level,
	}

	return &message
}

