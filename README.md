# logs and logos

---
## λόγος (logos) – Meaning "word," "speech," "reason," or "account."
#### (This is where "log" (as in a record or logbook) comes from.)

---


Common log interface for personal projects. Some features I wanted:
* Print function that always logs, regardless of level. 
* Convenience functions like Debug, Info, and Error. 
* No separate Debugf, Infof, Errorf functions.
* "With" statements for adding context.
* Ability to change log levels and names.
* Ability to change log formats.
* Ability to user a logger on its own, or set it as a default logger.

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
    log.Log(levels.Debug, "debug %s ...", "zap")
    log.Log(levels.Info, "info %s ...", "zap")
    log.Log(levels.Warn, "warn %s ...", "zap")
}
```

SLog:
```go
func main() {
    log = contrib.NewSLogLogger(levels.Debug, formats.Console, os.Stdout)
    log.Log(levels.Debug, "debug %s ...", "slog")
    log.Log(levels.Info, "info %s ...", "slog")
    log.Log(levels.Warn, "warn %s ...", "slog")
}
```
Logrus:
```go
func main() {
    log = contrib.NewLogrusLogger(levels.Debug, formats.Console, os.Stdout)
    log.Log(levels.Debug, "debug %s ...", "logrus")
    log.Log(levels.Info, "info %s ...", "logrus")
    log.Log(levels.Warn, "warn %s ...", "logrus")
}
```