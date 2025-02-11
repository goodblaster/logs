package main

import (
	"fmt"
	"os"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/colors"
	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
	"github.com/goodblaster/logs/pkg/contrib"
	"github.com/goodblaster/logs/pkg/logos"
)

func main() {
	log := logos.NewLogger(levels.Debug, formats.Console, os.Stdout)
	log.Debug("logos ...")

	logs.SetDefaultLogger(log)
	logs.Debug("logos as default logger ...")

	log.With("key", 9).Error("testing ...")
	log = log.With("key2", "two")
	log.Warn("warning ...")
	log = log.WithFields(map[string]interface{}{"key3": 3, "key4": "four"})
	log.Info("info ...")
	log = log.With("key5", map[string]any{
		"key6": 6,
	})
	log.Debug("debug ...")

	log = logos.NewLogger(levels.Error, formats.Console, os.Stdout)
	log.(*logos.Logger).LogFunc(levels.Debug, func() string {
		fmt.Println("log func is being called")
		return "log func called"
	})

	log = logos.NewLogger(levels.Error, formats.Console, os.Stdout)
	log.(*logos.Logger).LogFunc(levels.Error, func() string {
		fmt.Println("log func is being called")
		return "log func called"
	})

	log = contrib.NewSLogLogger(levels.Debug, formats.Console, os.Stdout)
	log.Log(levels.Debug, "debug slog ...")
	log.Log(levels.Info, "info slog ...")
	log.Log(levels.Warn, "warn slog ...")

	log = contrib.NewZapLogger(levels.Debug, formats.Console, os.Stdout)
	log.Log(levels.Debug, "debug zap ...")
	log.Log(levels.Info, "info zap ...")
	log.Log(levels.Warn, "warn zap ...")

	log = contrib.NewLogrusLogger(levels.Debug, formats.Console, os.Stdout)
	log.Log(levels.Debug, "debug logrus ...")
	
	const (
		LevelApple levels.Level = iota
		LevelBanana
		LevelCherry
	)

	levels.LevelNames = map[levels.Level]string{
		LevelApple:  "apple",
		LevelBanana: "banana",
		LevelCherry: "cherry",
	}

	levels.LevelColors = map[levels.Level]colors.TextColor{
		LevelApple:  colors.TextGreen,
		LevelBanana: colors.TextYellow,
		LevelCherry: colors.TextRed,
	}

	log = logos.NewLogger(LevelApple, formats.Console, os.Stdout)
	log.Log(LevelApple, "apple ...")
	log.Log(LevelBanana, "banana ...")
	log.Log(LevelCherry, "cherry ...")

	logs.SetDefaultLogger(log)
	logs.Log(LevelApple, "default apple ...")
}
