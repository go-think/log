package log

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/go-think/log/handler"
	"github.com/go-think/log/record"
)

func TestLog(t *testing.T) {
	Debug("log with Debug")
	Info("log with Info")
	Notice("log with Notice")
	Warn("log with Warn")
	Error("log with Error")
	Crit("log with Crit")
	Alert("log with Alert")
	Emerg("log with Emerg")
}

func TestLogWithFileHandler(t *testing.T) {

	filename := path.Join(os.TempDir(), "thinkgo.log")

	h := handler.NewFileHandler(filename, record.INFO)

	l := NewLogger("testing", record.INFO)
	l.PushHandler(h)

	filename = h.GetFilename()

	os.Remove(filename)

	message := "Log write to file"

	l.Debug(message)

	_, err := os.ReadFile(filename)
	if err == nil {
		t.Error("expected error")
	}

	h.SetLevel(record.DEBUG)
	l = NewLogger("testing", record.DEBUG)
	l.PushHandler(h)
	l.Debug(message)

	b, err := os.ReadFile(filename)
	if err != nil {
		t.Error(fmt.Errorf("read file %s error: %w", filename, err))
	}
	content := string(b)

	if !strings.Contains(content, message) {
		t.Error("test FileHandler error")
	}

}
