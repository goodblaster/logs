package formatters

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/goodblaster/logs"
	"github.com/stretchr/testify/assert"
)

func TestNewJsonFormatter(t *testing.T) {
	// Default config. Timestamp is close.
	cfg := DefaultConfig
	fmtr := NewJsonFormatter(cfg)
	line := fmtr.Format(logs.LevelInfo, "Test", nil)
	var m map[string]any //
	assert.NoError(t, json.Unmarshal([]byte(line), &m))
	then, err := time.ParseInLocation(DefaultTimestampFormat, m["timestamp"].(string), time.Local)
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now().Local(), then.Local(), time.Second)

	// Custom config. Timestamp is static.
	static := "2020-01-01T00:00:00"
	cfg = Config{
		Timestamp: func() string {
			return static
		},
	}
	fmtr = NewJsonFormatter(cfg)
	line = fmtr.Format(logs.LevelInfo, "Test", nil)
	assert.NoError(t, json.Unmarshal([]byte(line), &m))
	assert.Equal(t, static, m["timestamp"])

	// UTC
	cfg = Config{
		Timestamp: func() string {
			return time.Now().UTC().Format(DefaultTimestampFormat)
		},
	}
	fmtr = NewJsonFormatter(cfg)
	line = fmtr.Format(logs.LevelInfo, "Test", nil)
	assert.NoError(t, json.Unmarshal([]byte(line), &m))
	then, err = time.ParseInLocation(DefaultTimestampFormat, m["timestamp"].(string), time.UTC)
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now().UTC(), then.UTC(), time.Second)
}

func Test_jsonFormatter_Format(t *testing.T) {
	type params struct {
		cfg Config
	}
	type args struct {
		level  logs.Level
		msg    string
		fields Fields
	}
	tests := []struct {
		name     string
		params   params
		args     args
		contains map[string]any
	}{
		{
			name: "Msg only",
			params: params{
				cfg: DefaultConfig,
			},
			args: args{
				level:  logs.LevelInfo,
				msg:    "Test",
				fields: nil,
			},
			contains: map[string]any{
				"msg": "Test",
			},
		},
		{
			name: "Msg with fields",
			params: params{
				cfg: DefaultConfig,
			},
			args: args{
				level: logs.LevelInfo,
				msg:   "Test",
				fields: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			contains: map[string]any{
				"msg": "Test",
				"fields": map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := jsonFormatter{
				cfg: tt.params.cfg,
			}
			got := f.Format(tt.args.level, tt.args.msg, tt.args.fields)
			var m map[string]any
			assert.NoError(t, json.Unmarshal([]byte(got), &m))
			for key, value := range tt.contains {
				assert.EqualValues(t, m[key], value)
			}

		})
	}
}
