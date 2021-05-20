package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	log.Print("Logging in Go!")

	lg := zap.NewExample().Sugar()
	defer lg.Sync()

	lg.Info("fetch the url from some web site.", "data-dir", "dir-type")
	lg.Infof("%d %s", 10, "this is error.")

	var logger *zap.Logger
	logger, _ = zap.NewProduction()
	logger.Info("this is from prod logger")

}
