package main

import (
	"github.com/xuanghu/log"
	"time"
)

func main() {
	log.Root().SetHandler(log.NewFileLvlHandler("log", 1024, "debug"))

	for {
		time.Sleep(time.Second)
		log.Debug("print", "time", time.Now().Unix())
		log.Debug("asd", "time", time.Now().Unix())
		log.Debug("esf", "time", time.Now().Unix())
		log.SetKeywords("esf")
	}
}
