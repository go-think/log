package handler

import (
	"fmt"
	"os"

	"github.com/go-think/log/record"
)

type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = map[record.Level]brush{
	record.EMERGENCY: newBrush("1;41"), // Emergency          Red background
	record.ALERT:     newBrush("1;35"), // Alert              purple
	record.CRITICAL:  newBrush("1;34"), // Critical           blue
	record.ERROR:     newBrush("1;31"), // Error              red
	record.WARNING:   newBrush("1;33"), // Warn               yellow
	record.INFO:      newBrush("1;36"), // Informational      sky blue
	record.DEBUG:     newBrush("1;32"), // Debug              green
	record.NOTICE:    newBrush("1;32"), // Trace              green
}

type ConsoleHandler struct {
	Handler
	level record.Level

	bubble bool
}

func NewConsoleHandler(level record.Level) *ConsoleHandler {
	return &ConsoleHandler{
		level:  level,
		bubble: true,
	}
}

// IsHandling Checks whether the given record will be handled by this handler.
func (h *ConsoleHandler) IsHandling(r record.Record) bool {
	return r.Level >= h.level
}

// Handle Handles a record.
func (h *ConsoleHandler) Handle(r record.Record) (bool, error) {
	if !h.IsHandling(r) {
		return false, nil
	}

	r.Formatted = h.GetFormatter().Format(r)

	err := h.write(r)
	if err != nil {
		return false, err
	}

	return false == h.bubble, nil
}

func (h *ConsoleHandler) write(r record.Record) error {
	message := colors[r.Level](r.Formatted)
	_, err := os.Stdout.Write(append([]byte(message)))
	if err != nil {
		return fmt.Errorf("write console error: %w", err)
	}
	return nil
}
