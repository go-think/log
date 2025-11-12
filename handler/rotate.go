package handler

import (
	"github.com/go-think/log/record"
)

func NewRotateHandler(filename string, level record.Level) *FileHandler {
	// h.timedFilename = h.GetTimedFilename()
	fileHandler := NewFileHandler(filename, level)
	fileHandler.rotate = true

	return fileHandler
}
