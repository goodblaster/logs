package main

import (
	"os"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/pkg/xlog"
)

func main() {
	log := xlog.NewLogger(logs.LevelDebug, logs.FormatConsole, os.Stdout)
	log.With("key", 9).Error("testing ...")
	log = log.With("key2", "two")
	log.Warn("warning ...")
	log = log.WithFields(map[string]interface{}{"key3": 3, "key4": "four"})
	log.Info("info ...")
	log = log.With("key5", map[string]any{
		"key6": 6,
	})
	log.Debug("debug ...")
}
