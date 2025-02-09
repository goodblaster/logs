package main

import (
	"os"

	"github.com/goodblaster/logs/pkg/contrib"
	"github.com/goodblaster/logs/pkg/formats"
	"github.com/goodblaster/logs/pkg/levels"
)

func main() {
	log := contrib.NewZapLogger(levels.Debug, formats.JSON, os.Stdout)
	log.Print("testing ...")
}
