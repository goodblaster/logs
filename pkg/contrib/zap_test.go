package contrib

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/pkg/formats"
	"github.com/goodblaster/logs/pkg/levels"
	"github.com/stretchr/testify/assert"
)

func TestZapJSON(t *testing.T) {
	var logBuffer bytes.Buffer
	log := NewZapLogger(levels.Debug, formats.JSON, &logBuffer)

	// Define test cases
	tests := []struct {
		format   string
		args     []any
		expected string
		fields   map[string]any
	}{
		{"simple", nil, "simple", nil},
		{"simple %s", []any{"string"}, "simple string", nil},
		{"simple %d", []any{123}, "simple 123", nil},
		{"simple", nil, "simple", map[string]any{"key": "value", "key2": "value2"}},
	}

	for _, tt := range tests {
		logBuffer.Reset() // Clear buffer before each test case
		logger := logs.Logger(&(*log))
		for key, value := range tt.fields {
			logger = logger.With(key, value)
		}
		for k, v := range tt.fields {
			logger = logger.With(k, v)
		}

		funcMap := map[string]func(format string, args ...any){
			"print": logger.Print,
			"info":  logger.Info,
			"warn":  logger.Warn,
			"error": logger.Error,
			"debug": logger.Debug,
		}

		for level, logFunc := range funcMap {
			logFunc(tt.format, tt.args...)
			m := mapFromBuffer(&logBuffer)
			assert.Contains(t, m.String("level"), level, "Unexpected log level for: %v", tt)
			assert.Equal(t, tt.expected, m.String("msg"), "Unexpected log message for: %v", tt)
			for key, value := range tt.fields {
				assert.Equal(t, value, m[key], "Unexpected field value for: %v", tt)
			}
		}
	}
}

type Map map[string]any

func mapFromBuffer(buf *bytes.Buffer) Map {
	b := buf.Bytes()
	buf.Reset()

	m := make(Map)
	_ = json.Unmarshal(b, &m)
	return m
}

func (m Map) String(key string) string {
	if str, ok := m[key].(string); ok {
		return str
	}
	return ""
}

func (m Map) Int(key string) int {
	if i, ok := m[key].(int); ok {
		return i
	}
	return 0
}

func TestZapText(t *testing.T) {
	var logBuffer bytes.Buffer
	log := NewZapLogger(levels.Debug, formats.Text, &logBuffer)

	// Define test cases
	tests := []struct {
		format   string
		args     []any
		expected string
		fields   map[string]any
	}{
		{"simple", nil, "simple", nil},
		{"simple %s", []any{"string"}, "simple string", nil},
		{"simple %d", []any{123}, "simple 123", nil},
		{"simple", nil, "simple", map[string]any{"key": "value", "key2": "value2"}},
	}

	for _, tt := range tests {
		logBuffer.Reset() // Clear buffer before each test case
		logger := logs.Logger(&(*log))
		for key, value := range tt.fields {
			logger = logger.With(key, value)
		}
		for k, v := range tt.fields {
			logger = logger.With(k, v)
		}

		funcMap := map[string]func(format string, args ...any){
			"print": logger.Print,
			"info":  logger.Info,
			"warn":  logger.Warn,
			"error": logger.Error,
			"debug": logger.Debug,
		}

		for level, logFunc := range funcMap {
			logFunc(tt.format, tt.args...)
			f := fieldsFromBuffer(&logBuffer)
			assert.Contains(t, f.String(1), level, "Unexpected log level for: %v", tt)
			assert.Equal(t, tt.expected, f.String(2), "Unexpected log message for: %v", tt)
			m := make(map[string]any)
			_ = json.Unmarshal([]byte(f.String(3)), &m)
			for key, value := range tt.fields {
				assert.Equal(t, value, m[key], "Unexpected field value for: %v", tt)
			}
		}
	}
}

type Fields []string

func (f Fields) String(i int) string {
	if i < len(f) {
		return f[i]
	}
	return ""
}

func fieldsFromBuffer(buf *bytes.Buffer) Fields {
	b := buf.Bytes()
	buf.Reset()

	var f Fields
	f = strings.Split(string(b), "\t")
	for i := range f {
		f[i] = strings.TrimSpace(f[i])
	}
	return f
}
