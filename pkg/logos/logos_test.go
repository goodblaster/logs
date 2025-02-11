package logos

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
	"github.com/stretchr/testify/assert"
)

func TestLogos_ConvenienceFunctions(t *testing.T) {
	buf := &bytes.Buffer{}

	// As debug, all logs should be printed
	log := NewLogger(levels.Debug, formats.JSON, buf)
	log.Debug("logos")
	m := Map(buf)
	assert.Equal(t, "logos", m["msg"])
	assert.Equal(t, "debug", m["level"])
	log.Info("logos")
	assert.Equal(t, "info", Map(buf)["level"])
	log.Warn("logos")
	assert.Equal(t, "warn", Map(buf)["level"])
	log.Error("logos")
	assert.Equal(t, "error", Map(buf)["level"])
	log.Print("logos")
	assert.Equal(t, "print", Map(buf)["level"])
}

func TestLogos_Levels(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)

	for _, level := range []levels.Level{levels.Debug, levels.Info, levels.Warn, levels.Error, levels.Print} {
		log.Log(level, "logos")
		assert.Equal(t, level.String(), Map(buf)["level"])
	}

	// Change the level to error, only error and print logs should be printed.
	log.SetLevel(levels.Error)
	log.Debug("logos")
	assert.Empty(t, buf.String())
	log.Info("logos")
	assert.Empty(t, buf.String())
	log.Warn("logos")
	assert.Empty(t, buf.String())
	log.Error("logos")
	assert.Equal(t, "error", Map(buf)["level"])
	log.Print("logos")
	assert.Equal(t, "print", Map(buf)["level"])
}

func TestLogos_With(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)

	log.With("key", "value").Log(levels.Debug, "logos")
	assert.Equal(t, "value", MapField(Map(buf), "key"))
}

func TestLogos_WithFields(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)

	log.WithFields(map[string]any{"key": "value"}).Log(levels.Debug, "logos")
	assert.Equal(t, "value", MapField(Map(buf), "key"))
}

func TestLogos_LogFunc(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)

	log.LogFunc(levels.Debug, func() string {
		return "logos"
	})
	assert.Equal(t, "logos", Map(buf)["msg"])
}

func TestLogos_SetLevel(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)

	log.SetLevel(levels.Error)
	log.Debug("logos")
	assert.Empty(t, buf.String())
	log.Error("logos")
	assert.Equal(t, "error", Map(buf)["level"])
}

func TestLogos_CustomLevels(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)

	const (
		LevelApple levels.Level = iota + 100
		LevelBanana
		LevelCherry
	)

	oldLevels := map[levels.Level]string{}
	for level, name := range levels.LevelNames {
		oldLevels[level] = name
	}

	defer func() {
		levels.LevelNames = oldLevels
	}()

	levels.LevelNames = map[levels.Level]string{
		LevelApple:  "apple",
		LevelBanana: "banana",
		LevelCherry: "cherry",
	}

	log.Log(LevelApple, "apple")
	assert.Equal(t, "apple", Map(buf)["level"])
	log.Log(LevelBanana, "banana")
	assert.Equal(t, "banana", Map(buf)["level"])
	log.Log(LevelCherry, "cherry")
	assert.Equal(t, "cherry", Map(buf)["level"])

	log.SetLevel(LevelBanana)
	log.Log(LevelApple, "apple")
	assert.Empty(t, buf.String())
	log.Log(LevelBanana, "banana")
	assert.Equal(t, "banana", Map(buf)["level"])
	log.Log(LevelCherry, "cherry")
	assert.Equal(t, "cherry", Map(buf)["level"])
}

func TestLogos_DefaultLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)
	logs.SetDefaultLogger(log)

	logs.Log(levels.Debug, "logos")
	assert.Equal(t, "debug", Map(buf)["level"])
}

func TestLogos_SubLoggers(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogger(levels.Debug, formats.JSON, buf)

	subLog := log.With("key", "value")
	sublog2 := subLog.With("key2", "value2")
	sublog3 := sublog2.With("key3", "value3")

	sublog3.Log(levels.Debug, "logos")
	m := Map(buf)
	assert.Equal(t, "value", MapField(m, "key"))
	assert.Equal(t, "value2", MapField(m, "key2"))
	assert.Equal(t, "value3", MapField(m, "key3"))

	sublog2.Log(levels.Debug, "logos")
	m = Map(buf)
	assert.Equal(t, "value", MapField(m, "key"))
	assert.Equal(t, "value2", MapField(m, "key2"))
	assert.Empty(t, MapField(m, "key3"))

	subLog.Log(levels.Debug, "logos")
	m = Map(buf)
	assert.Equal(t, "value", MapField(m, "key"))
	assert.Empty(t, MapField(m, "key2"))
	assert.Empty(t, MapField(m, "key3"))
}

func Map(buf *bytes.Buffer) map[string]any {
	m := make(map[string]any)
	_ = json.Unmarshal(buf.Bytes(), &m)
	buf.Reset()
	return m
}

func MapField(m map[string]any, key string) any {
	return m["fields"].(map[string]any)[key]
}
