# logs and logos

---
## λόγος (logos) – Meaning "word," "speech," "reason," or "account."
#### (This is where "log" (as in a record or logbook) comes from.)

---


Common log interface for personal projects. I wanted a Print function that always logs. 
I wanted convenience functions like Debug, Info, and Error. 
I wanted them to handle format/args instead of requiring separate Debugf, Infof, and Errorf functions.
I wanted With statements to add context to logs.

Logos is my own implementation that was born out of this exercise.
I found much of the configuration for the existing loggers tedious and unnecessarily heavy-weight.

It is simple to use and modify.
Getting started:
```go
    log := logos.NewLogger(levels.Debug, formats.Console, os.Stdout)
    log.Debug("logos ...")
```
You can set it as a default logger so you don't have to pass the logger around.
```go
    logs.SetDefaultLogger(log)
    logs.Debug("default ...")
```

If you want to change log levels, and names, do whatever you like. Example:

```go
func main() {
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

	log := logos.NewLogger(LevelApple, formats.Console, os.Stdout)
	log.Log(LevelApple, "apple ...")
	log.Log(LevelBanana, "banana ...")
	log.Log(LevelCherry, "cherry ...")

	// Or set as default.
	logs.SetDefaultLogger(log)
	logs.Log(LevelApple, "default apple ...")
}
```

If you want to use the Zap logger:
```go
func main() {
    log := contrib.NewZapLogger(levels.Debug, formats.Console, os.Stdout)
    log.Log(levels.Debug, "debug zap ...")
    log.Log(levels.Info, "info zap ...")
    log.Log(levels.Warn, "warn zap ...")
}
```

SLog:
```go
func main() {
    log = contrib.NewSLogLogger(levels.Debug, formats.Console, os.Stdout)
    log.Log(levels.Debug, "debug slog ...")
    log.Log(levels.Info, "info slog ...")
    log.Log(levels.Warn, "warn slog ...")
}
```
Logrus:
```go
func main() {
    log = contrib.NewLogrusLogger(levels.Debug, formats.Console, os.Stdout)
    log.Log(levels.Debug, "debug logrus ...")
    log.Log(levels.Info, "info logrus ...")
    log.Log(levels.Warn, "warn logrus ...")
}
```