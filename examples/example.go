package main

import (
	"errors"
	log "github.com/Ifkarsyah/golog"
)

func main() {
	log.Success("ye")
	log.Info("this is just info")
	log.Warn("why here")
	log.Debug(errors.New("this is debug example"))
	log.Fatal(errors.New("stop"))
}
