package handler

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-think/log/record"
)

type FileHandler struct {
	Handler
	level    record.Level
	bubble   bool
	filename string

	filenameFormat string
	dateFormat     string
	timedFilename  string

	rotate bool
}

func NewFileHandler(filename string, level record.Level) *FileHandler {
	h := &FileHandler{
		level:          level,
		bubble:         true,
		filename:       filename,
		filenameFormat: "{filename}-{date}",
		dateFormat:     "2006-01-02",
	}
	// h.timedFilename = h.GetTimedFilename()
	return h
}

// IsHandling Checks whether the given record will be handled by this handler.
func (h *FileHandler) IsHandling(r record.Record) bool {
	return r.Level >= h.level
}

// Handle Handles a record.
func (h *FileHandler) Handle(r record.Record) (bool, error) {
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

// SetLevel Sets minimum logging level at which this handler will be triggered.
func (h *FileHandler) SetLevel(level record.Level) {
	h.level = level
}

func (h *FileHandler) write(r record.Record) error {
	file, err := os.OpenFile(h.GetFilename(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open file %s error: %w", h.filename, err)
	}
	defer file.Close()
	_, err = file.Write([]byte(r.Formatted))
	if err != nil {
		return fmt.Errorf("write file %s error: %w", h.filename, err)
	}
	return nil
}

// GetFilename Gets the filename.
func (h *FileHandler) GetFilename() string {
	if !h.rotate {
		return h.filename
	}

	return h.GetTimedFilename()
}

// GetTimedFilename Gets the timed filename.
func (h *FileHandler) GetTimedFilename() string {
	dirname := path.Dir(h.filename)
	filename := path.Base(h.filename)
	fileExt := path.Ext(h.filename)
	filename = strings.TrimSuffix(filename, fileExt)

	timedFilename := strings.Replace(path.Join(dirname, h.filenameFormat), "{filename}", filename, -1)
	timedFilename = strings.Replace(timedFilename, "{date}", time.Now().Local().Format(h.dateFormat), -1)

	if len(fileExt) > 0 {
		timedFilename = timedFilename + fileExt
	}

	return timedFilename
}
