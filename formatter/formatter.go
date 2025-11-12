package formatter

import "github.com/go-think/log/record"

type Formatter interface {
	Format(r record.Record) string
	FormatBatch(rs []record.Record) string
}
