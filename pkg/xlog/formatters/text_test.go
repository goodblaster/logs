package formatters

import (
	"strings"
	"testing"
	"time"

	"github.com/goodblaster/logs"
	"github.com/stretchr/testify/assert"
)

func TestNewTextFormatter(t *testing.T) {
	cfg := DefaultConfig
	fmtr := NewTextFormatter(cfg)
	assert.NotNil(t, fmtr)

	// Default config. Timestamp is close.
	cfg = DefaultConfig
	fmtr = NewTextFormatter(cfg)
	line := fmtr.Format(logs.LevelInfo, "Test", nil)
	assert.Equal(t, "info", strings.Fields(line)[1])
	assert.Equal(t, "Test", strings.Fields(line)[2])
	then, err := time.ParseInLocation(DefaultTimestampFormat, strings.Fields(line)[0], time.Local)
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now().Local(), then.Local(), time.Second)

	// static timestamp
	static := "2020-01-01T00:00:00"
	cfg = Config{
		Timestamp: func() string {
			return static
		},
	}
	fmtr = NewTextFormatter(cfg)
	line = fmtr.Format(logs.LevelPrint, "Test", nil)
	assert.Equal(t, static, strings.Fields(line)[0])
	assert.Equal(t, "print", strings.Fields(line)[1])
	assert.Equal(t, "Test", strings.Fields(line)[2])

	// UTC
	cfg = Config{
		Timestamp: func() string {
			return time.Now().UTC().Format(DefaultTimestampFormat)
		},
	}
	fmtr = NewTextFormatter(cfg)
	line = fmtr.Format(logs.LevelInfo, "Test", nil)
	assert.Equal(t, "info", strings.Fields(line)[1])
	assert.Equal(t, "Test", strings.Fields(line)[2])
	then, err = time.ParseInLocation(DefaultTimestampFormat, strings.Fields(line)[0], time.UTC)
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now().UTC(), then.UTC(), time.Second)

	// With some fields.
	line = fmtr.Format(logs.LevelInfo, "Test", Fields{"key": "value"})
	assert.Equal(t, "info", strings.Fields(line)[1])
	assert.Equal(t, "key=value", strings.Fields(line)[2])
	assert.Equal(t, "Test", strings.Fields(line)[3])
}

func Test_textFormatter_Format(t *testing.T) {
	type fields struct {
		cfg Config
	}
	type args struct {
		level  logs.Level
		msg    string
		fields Fields
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := textFormatter{
				cfg: tt.fields.cfg,
			}
			assert.Equalf(t, tt.want, f.Format(tt.args.level, tt.args.msg, tt.args.fields), "Format(%v, %v, %v)", tt.args.level, tt.args.msg, tt.args.fields)
		})
	}
}
